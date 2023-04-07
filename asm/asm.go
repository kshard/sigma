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

package asm

import (
	"bytes"
	"fmt"

	"github.com/kshard/sigma/vm"
)

//
// The file defines byte-code (assembler) for expressing on σ-calculus
//

// Context for linker
type Context struct {
	Facts map[string]vm.Generator
}

func (ctx *Context) Add(id string, gen vm.Generator) *Context {
	ctx.Facts[id] = gen
	return ctx
}

func NewContext() *Context {
	return &Context{
		Facts: make(map[string]vm.Generator),
	}
}

type Stream interface {
	Link(*Context) vm.Stream
}

type Generator struct {
	Name string
	Addr []vm.Addr
}

func (gen *Generator) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(gen.Name)
	buffer.WriteString("(")
	for _, addr := range gen.Addr[:len(gen.Addr)-1] {
		buffer.WriteString(fmt.Sprintf("%v", addr))
		buffer.WriteString(",")
	}
	buffer.WriteString(fmt.Sprintf("%v", gen.Addr[len(gen.Addr)-1]))
	buffer.WriteString(")")
	return buffer.String()
}

func (gen *Generator) Link(ctx *Context) vm.Stream {
	fact, exists := ctx.Facts[gen.Name]
	if !exists {
		panic(fmt.Errorf("unknown %v", gen.Name))
	}

	return fact(gen.Addr)
}

type Horn struct {
	Body []Stream
}

func (horn *Horn) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, f := range horn.Body {
		if i == 0 {
			buffer.WriteString(" ")
		}
		buffer.WriteString(fmt.Sprintf("%v", f))
		buffer.WriteString(" ")
	}
	buffer.WriteString("]")
	return buffer.String()
}

func (horn *Horn) Link(ctx *Context) vm.Stream {
	seq := make([]vm.Stream, len(horn.Body))
	for i, f := range horn.Body {
		seq[i] = f.Link(ctx)
	}

	return vm.Horn(seq...)
}
