package compile

import (
	"fmt"

	"github.com/0xdbf/sigma/ast"
	"github.com/0xdbf/sigma/internal/vm"
)

type Context struct {
	Rules   map[string]vm.Stream
	Facts   map[string]func([]vm.Addr) vm.Stream
	Memory  map[string]vm.Addr
	Literal map[vm.Addr]any
}

func New() *Context {
	return &Context{
		Rules:   make(map[string]vm.Stream),
		Facts:   make(map[string]func([]vm.Addr) vm.Stream),
		Memory:  make(map[string]vm.Addr),
		Literal: make(map[vm.Addr]any),
	}
}

func (ctx *Context) Heap() *vm.Heap {
	heap := make(vm.Heap, len(ctx.Memory))
	for addr, val := range ctx.Literal {
		heap.Put(addr, val)
	}
	return &heap
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
			if err := ctx.compileHorn(horn); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ctx *Context) compileHorn(horn *ast.Horn) error {
	seq := []vm.Stream{}

	for _, imply := range horn.Body {
		stream, exists := ctx.Rules[imply.Name]
		if exists {
			_, err := ctx.compileTerms(imply)
			if err != nil {
				return err
			}
			// TODO: wrap invocation of horn
			seq = append(seq, stream)
			continue
		}

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

	ctx.Rules[horn.Head.Name] = vm.Horn(seq...)
	return nil
}

//
func (ctx *Context) compileTerms(obj *ast.Imply) ([]vm.Addr, error) {
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

	return addr | (1 << 31)
}

func (ctx *Context) literal(name string, value *any) vm.Addr {
	addr, exists := ctx.Memory[name]

	if !exists {
		addr = vm.Addr(len(ctx.Memory))
		ctx.Memory[name] = addr
		ctx.Literal[addr] = value
		return addr | (1 << 31)
	}

	return addr | (1 << 31)
}
