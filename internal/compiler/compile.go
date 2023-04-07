/*

  Sigma Virtual Machine
  Copyright (C) 2016  Dmitry Kolesnikov

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as published
  by the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/

package compiler

import (
	"fmt"

	"github.com/kshard/sigma/asm"
	"github.com/kshard/sigma/ast"

	"github.com/kshard/sigma/vm"
)

type Compiler func(*Context, ast.Terms) asm.Stream

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

// Create new instance of compiler
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

// Assemble the goal of σ-application into executable object.
// It returns virtual machine, target relation and stream itself
func (ctx *Context) Assemble(goal string) (*vm.VM, []vm.Addr, asm.Stream) {
	compiler := ctx.Rules[goal]
	stream := compiler(ctx, nil)

	head := ctx.Signs[goal]
	addr := make([]vm.Addr, len(head.Terms))
	for i, term := range head.Terms {
		addr[i] = ctx.Heap[term.Name]
	}

	vmm := vm.New(len(ctx.Heap))
	for addr, val := range ctx.Const {
		vmm.Heap.Put(addr, val)
	}

	return vmm, addr, stream
}

func (ctx *Context) Compile(rules ast.Rules) error {
	for _, rule := range rules {
		switch horn := rule.(type) {
		case *ast.Fact:
			ctx.Facts[horn.Stream.Name] = func(_ []vm.Addr) vm.Stream { return nil }
		case *ast.Horn:
			ctx.Signs[horn.Head.Name] = horn.Head
			ctx.Rules[horn.Head.Name] = ctx.horn(horn)
		}
	}

	return nil
}

func (ctx *Context) horn(horn *ast.Horn) Compiler {
	return func(ctx *Context, args ast.Terms) asm.Stream {
		refs := ctx.head(horn.Head, args)

		body := []asm.Stream{}
		for _, imply := range horn.Body {
			seq := ctx.stmt(imply.Terms, refs)

			rule, exists := ctx.Rules[imply.Name]
			if exists {
				body = append(body, rule(ctx, seq))
				continue
			}

			_, exists = ctx.Facts[imply.Name]
			if exists {
				addr := make([]vm.Addr, len(seq))
				for i, term := range seq {
					addr[i] = ctx.alloc(term)
				}
				// body = append(body, fact(addr))
				body = append(body, &asm.Generator{Name: imply.Name, Addr: addr})

				continue
			}
		}

		return &asm.Horn{Body: body} //vm.Horn(body...)
	}
}

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
