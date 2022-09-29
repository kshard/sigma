/*

  Copyright 2016 Dmitry Kolesnikov, All Rights Reserved

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package compile

import (
	"fmt"

	"github.com/0xdbf/sigma/ast"

	"github.com/0xdbf/sigma/vm"
)

type Compiler func(*Context, ast.Terms) vm.Stream

type Rules map[string]Compiler

type Refs map[string]*ast.Term

type Heap map[string]vm.Addr

type Context struct {
	Signs map[string]*ast.Head
	Rules Rules
	Facts map[string]vm.Generator
	Heap  Heap
	Const map[vm.Addr]any
	Index int
}

func New() *Context {
	return &Context{
		Signs: make(map[string]*ast.Head),
		Rules: Rules{},
		Facts: make(map[string]vm.Generator),
		Heap:  Heap{},
		Const: make(map[vm.Addr]any),
		Index: 0,
	}
}

func (ctx *Context) Reader(goal string) *vm.Reader {
	compiler := ctx.Rules[goal]
	stream := compiler(ctx, nil)

	vmm := vm.New(len(ctx.Heap))
	for addr, val := range ctx.Const {
		vmm.Heap.Put(addr, val)
	}

	head := ctx.Signs[goal]
	addr := make([]vm.Addr, len(head.Terms))
	for i, term := range head.Terms {
		addr[i] = ctx.Heap[term.Name]
	}

	return vmm.Stream(addr, stream)
}

func (ctx *Context) Compile(rules ast.Rules) error {
	for _, rule := range rules {
		switch horn := rule.(type) {
		case *ast.Fact:
			ctx.Facts[horn.Stream.Name] = horn.Generator(horn.Stream.Terms)
		case *ast.Horn:
			ctx.Signs[horn.Head.Name] = horn.Head
			ctx.Rules[horn.Head.Name] = ctx.horn(horn)
		}
	}

	return nil
}

func (ctx *Context) horn(horn *ast.Horn) Compiler {
	return func(ctx *Context, args ast.Terms) vm.Stream {
		// fmt.Printf("%s%v => %v\n", horn.Head.Name, args, horn.Head.Terms)
		refs := ctx.head(horn.Head, args)
		// fmt.Println(refs)

		body := []vm.Stream{}
		for _, imply := range horn.Body {
			seq := ctx.stmt(imply.Terms, refs)
			// fmt.Printf("%s%v\n", imply.Name, seq)

			rule, exists := ctx.Rules[imply.Name]
			if exists {
				body = append(body, rule(ctx, seq))
				continue
			}

			fact, exists := ctx.Facts[imply.Name]
			if exists {
				addr := make([]vm.Addr, len(seq))
				for i, term := range seq {
					addr[i] = ctx.alloc(term)
				}
				// fmt.Printf("==> %s%v\n", imply.Name, addr)
				body = append(body, fact(addr))

				continue
			}

			// Generator
			// fmt.Printf("%s%v\n", imply.Name, imply.Terms)
		}

		return vm.Horn(body...)
	}
}

//
func (ctx *Context) head(head *ast.Head, args ast.Terms) Refs {
	terms := Refs{}

	for i, t := range head.Terms {
		if args != nil {
			terms[t.Name] = args[i]
		} else {
			terms[t.Name] = t
		}
	}

	return terms
}

//
func (ctx *Context) stmt(args ast.Terms, refs Refs) ast.Terms {
	seq := make(ast.Terms, len(args))

	for i, arg := range args {
		if r, has := refs[arg.Name]; has {
			seq[i] = r
		} else {
			if arg.Value == nil {
				name := fmt.Sprintf("%s%d", arg.Name, ctx.Index)

				refs[arg.Name] = &ast.Term{Name: name}
				seq[i] = refs[arg.Name]

				ctx.Index++
			} else {
				seq[i] = arg
			}
		}
	}

	return seq
}

func (ctx *Context) alloc(term *ast.Term) vm.Addr {
	addr, exists := ctx.Heap[term.Name]

	if !exists {
		addr = vm.Addr(len(ctx.Heap))

		if term.Value != nil {
			ctx.Heap[term.Name] = addr
			ctx.Const[addr] = &term.Value
			// fmt.Printf("==> %s = %v\n", term.Name, addr.ReadOnly())
			return addr.ReadOnly()
		}

		ctx.Heap[term.Name] = addr
		// fmt.Printf("==> %s = %v\n", term.Name, addr)

		return addr
	}

	// fmt.Printf("==> %s = %v\n", term.Name, addr.ReadOnly())
	return addr.ReadOnly()
}

/*

type Rule struct {
	// Head []
	Body Stream
}

type Context struct {
	//
	Rules   map[string]Rule
	Literal map[vm.Addr]any
	Memory  Heap
	Stack   []string

	//
	Facts map[string]func([]vm.Addr) vm.Stream
	// Memory map[string]vm.Addr
}

func New() *Context {
	return &Context{
		Rules:   make(map[string]Rule),
		Literal: make(map[vm.Addr]any),
		Memory:  make(Heap),
		Stack:   make([]string, 0),

		Facts: make(map[string]func([]vm.Addr) vm.Stream),
		// Memory: make(map[string]vm.Addr),
	}
}


func (ctx *Context) Compile(rules ast.Rules) error {
	// ctx.compileFacts(rules)
	return ctx.compileRules(rules)
}

//
// func (ctx *Context) compileFacts(rules []ast.Rule) {
// 	for _, rule := range rules {
// 		switch fact := rule.(type) {
// 		case *ast.Fact:
// 			ctx.Facts[fact.Name] = fact
// 		}
// 	}
// }

//
func (ctx *Context) compileRules(rules ast.Rules) error {
	for _, rule := range rules {
		switch horn := rule.(type) {
		case *ast.Horn:
			ctx.Rules[horn.Head.Name] = Rule{
				Body: Horn(horn),
			}
			// if err := ctx.compileHorn(horn); err != nil {
			// 	return err
			// }
		}
	}
	return nil
}

func (ctx *Context) compileHorn(horn *ast.Horn) error {
	seq := []vm.Stream{}

	// Note always writable
	// addr, _ := ctx.compileHead(horn.Head)

	for _, imply := range horn.Body {
		// TODO: refactor based of stack trace
		// stream, exists := ctx.Rules[imply.Name]
		// if exists {
		// 	terms, err := ctx.compileTerms(imply)
		// 	fmt.Printf("%s%v\n", imply.Name, terms)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	// seq = append(seq, vm.Call(terms, stream.Head, stream.Body))
		// 	seq = append(seq, stream.Body)
		// 	continue
		// }

		facts, exists := ctx.Facts[imply.Name]
		if exists {
			terms, err := ctx.compileTerms(imply)
			if err != nil {
				return err
			}
			seq = append(seq, facts(terms))
			continue
		}

		return fmt.Errorf("the rule %s is not defined", imply.Name)
	}

	// addr, _ := ctx.compileTerms(horn.Head)
	// xxx := make([]vm.Addr, len(addr))
	// for i, x := range addr {
	// 	xxx[i] = vm.Addr(x.Value())
	// }

	// ctx.Rules[horn.Head.Name] = Rule{
	// 	// Head: addr,
	// 	Body: vm.Horn(seq...),
	// }

	return nil
}

// func (ctx *Context) compileHead(head *ast.Imply) ([]vm.Addr, error) {
// 	seq := []vm.Addr{}

// 	for _, t := range head.Terms {
// 		switch tt := t.(type) {
// 		case *ast.Var:
// 			seq = append(seq, ctx.allochead(tt.Name))
// 		default:
// 			return nil, fmt.Errorf("xxx")
// 		}
// 	}

// 	return seq, nil
// }

//
func (ctx *Context) compileTerms(obj *ast.Imply) ([]vm.Addr, error) {
	// TODO: annotate variable terms with rule ID
	//       variables are scoped to conjunctive (horn) query
	//

	seq := []vm.Addr{}

	for _, t := range obj.Terms {
		switch tt := t.(type) {
		case *ast.Var:
			seq = append(seq, ctx.alloc(tt.Name))
		case *ast.Lit:
			seq = append(seq, ctx.literal(tt.ID, &tt.Value))
		}
	}

	return seq, nil
}

// allocate memory address to named object
func (ctx *Context) alloc(name string) vm.Addr {
	addr, exists := ctx.Memory[name]

	if !exists {
		addr = vm.Addr(len(ctx.Memory))
		ctx.Memory[name] = addr
		return addr
	}

	return addr.ReadOnly()
}

// allocate memory address to named object
// func (ctx *Context) allochead(name string) vm.Addr {
// 	addr, exists := ctx.Memory[name]

// 	if !exists {
// 		addr = vm.Addr(len(ctx.Memory)).Pointer()
// 		ctx.Memory[name] = addr
// 		fmt.Printf("he var %s = %v\n", name, addr)
// 		return addr
// 	}

// 	return addr //.ReadOnly()
// }

func (ctx *Context) literal(name string, value *any) vm.Addr {
	addr, exists := ctx.Memory[name]

	if !exists {
		addr = vm.Addr(len(ctx.Memory))
		ctx.Memory[name] = addr
		ctx.Literal[addr] = value
		return addr.ReadOnly()
	}

	return addr.ReadOnly()
}

*/
