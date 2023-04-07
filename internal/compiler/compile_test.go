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

package compiler_test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"testing"

	"github.com/kshard/sigma/asm"
	"github.com/kshard/sigma/ast"
	"github.com/kshard/sigma/internal/compiler"
	"github.com/kshard/sigma/internal/gen"
	"github.com/kshard/sigma/vm"
)

func TestX(t *testing.T) {
	e := ast.Rules{
		&ast.Fact{
			Stream: &ast.Imply{Name: "f", Terms: ast.Terms{{Name: "s"}, {Name: "p"}, {Name: "o"}}},
			// Generator: gen.FactsIMDB,
		},

		&ast.Horn{
			Head: &ast.Head{Name: "a", Terms: ast.Terms{{Name: "movie"}, {Name: "cast"}}},
			Body: ast.Implies{
				{Name: "f", Terms: ast.Terms{
					{Name: "m"},
					{Name: "t1", Value: "title"},
					{Name: "movie"},
				}},
				{Name: "f", Terms: ast.Terms{
					{Name: "m"},
					{Name: "c1", Value: "cast"},
					{Name: "cast"},
				}},
			},
		},

		&ast.Horn{
			Head: &ast.Head{Name: "h", Terms: ast.Terms{{Name: "name"}, {Name: "name1"}}},
			Body: ast.Implies{
				{Name: "a", Terms: ast.Terms{
					{Name: "t2", Value: "Lethal Weapon"},
					{Name: "p"},
				}},
				{Name: "f", Terms: ast.Terms{
					{Name: "p"},
					{Name: "n1", Value: "name"},
					{Name: "name"},
				}},
				{Name: "a", Terms: ast.Terms{
					{Name: "t3", Value: "Mad Max"},
					{Name: "s"},
				}},
				{Name: "f", Terms: ast.Terms{
					{Name: "s"},
					{Name: "n1", Value: "name"},
					{Name: "name1"},
				}},
			},
		},
	}

	c := compiler.New()
	c.Compile(e)

	vmm, addr, reader := c.Assemble("h")
	fmt.Println(*vmm, addr, reader)

	// raw, _ := json.Marshal(reader)
	// fmt.Println(string(raw))

	// var asmx asm.Horn
	// err := json.Unmarshal(raw, &asmx)
	// if err != nil {
	// 	panic(err)
	// }

	var buf bytes.Buffer
	gob.Register(&asm.Generator{})
	gob.Register(&asm.Horn{})
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(reader); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())

	var asmx asm.Horn
	dec := gob.NewDecoder(&buf)
	if err := dec.Decode(&asmx); err != nil {
		log.Fatal(err)
	}

	ctx := &asm.Context{
		Facts: map[string]vm.Generator{
			"f": func(addr []vm.Addr) vm.Stream {
				return gen.NewSubQ(addr, gen.IMDB())
			},
		},
	}

	value := make([]any, 2)
	xreader := vmm.Stream(addr, asmx.Link(ctx))

	for {
		if err := xreader.Read(value); err != nil {
			break
		}
		fmt.Println(value)
	}
}

/*
func TestCompile(t *testing.T) {
	e := ast.Rules{
		&ast.Horn{
			Head: &ast.Imply{Name: "h", Terms: ast.Terms{&ast.Var{"name"}}},
			Body: []*ast.Imply{
				{Name: "f", Terms: ast.Terms{
					&ast.Var{"m"},
					&ast.Lit{"t1", "title"},
					&ast.Lit{"t2", "Lethal Weapon"},
				}},
				{Name: "f", Terms: ast.Terms{
					&ast.Var{"m"},
					&ast.Lit{"c1", "cast"},
					&ast.Var{"p"},
				}},
				{Name: "f", Terms: ast.Terms{
					&ast.Var{"p"},
					&ast.Lit{"n1", "name"},
					&ast.Var{"name"},
				}},
			},
		},
	}

	c := compile.New()
	c.Facts["f"] = func(addr []vm.Addr) vm.Stream {
		return gen.NewSubQ(addr, gen.IMDB())
	}

	c.Compile(e)
	h, s := c.Heap("h")

	vm.Debug(s, h)
}

func TestXxx(t *testing.T) {
	e := ast.Rules{
		&ast.Horn{
			Head: &ast.Imply{Name: "a", Terms: ast.Terms{&ast.Var{"movie"}, &ast.Var{"cast"}}},
			Body: []*ast.Imply{
				{Name: "f", Terms: ast.Terms{
					&ast.Var{"m"},
					&ast.Lit{"t1", "title"},
					&ast.Var{"movie"},
				}},
				{Name: "f", Terms: ast.Terms{
					&ast.Var{"m"},
					&ast.Lit{"c1", "cast"},
					&ast.Var{"cast"},
				}},
			},
		},

		// &ast.Horn{
		// 	Head: &ast.Imply{Name: "b", Terms: ast.Terms{&ast.Var{"movie"}, &ast.Var{"cast"}}},
		// 	Body: []*ast.Imply{
		// 		{Name: "z", Terms: ast.Terms{
		// 			&ast.Var{"bm"},
		// 			&ast.Lit{"t1", "title"},
		// 			&ast.Var{"movie"},
		// 		}},
		// 		{Name: "z", Terms: ast.Terms{
		// 			&ast.Var{"bm"},
		// 			&ast.Lit{"c1", "cast"},
		// 			&ast.Var{"cast"},
		// 		}},
		// 	},
		// },

		&ast.Horn{
			Head: &ast.Imply{Name: "h", Terms: ast.Terms{&ast.Var{"name"}}},
			Body: []*ast.Imply{
				{Name: "a", Terms: ast.Terms{
					&ast.Lit{"t2", "Lethal Weapon"},
					&ast.Var{"p"},
				}},
				// {Name: "b", Terms: ast.Terms{
				// 	&ast.Lit{"t3", "Mad Max"},
				// 	&ast.Var{"s"},
				// }},
				{Name: "f", Terms: ast.Terms{
					&ast.Var{"p"},
					&ast.Lit{"n1", "name"},
					&ast.Var{"name"},
				}},
				// {Name: "f", Terms: ast.Terms{
				// 	&ast.Var{"s"},
				// 	&ast.Lit{"n1", "name"},
				// 	&ast.Var{"na"},
				// }},
			},
		},
	}
	c := compile.New()
	c.Facts["f"] = func(addr []vm.Addr) vm.Stream {
		return gen.NewSubQ(addr, gen.IMDB())
	}
	// c.Facts["z"] = func(addr []vm.Addr) vm.Stream {
	// 	return gen.NewSubQ(addr, gen.IMDB())
	// }

	c.Compile(e)
	h, s := c.Heap("h")

	vm.Debug(s, h)
}
*/

func BenchmarkTx(bb *testing.B) {
	e := ast.Rules{
		&ast.Fact{
			Stream: &ast.Imply{Name: "f", Terms: ast.Terms{{Name: "s"}, {Name: "p"}, {Name: "o"}}},
		},

		&ast.Horn{
			Head: &ast.Head{Name: "h", Terms: ast.Terms{{Name: "name"}}},
			Body: []*ast.Imply{
				{Name: "f", Terms: ast.Terms{
					{Name: "m"},
					{Name: "t1", Value: "title"},
					{Name: "t2", Value: "Lethal Weapon"},
				}},
				{Name: "f", Terms: ast.Terms{
					{Name: "m"},
					{Name: "c1", Value: "cast"},
					{Name: "p"},
				}},
				{Name: "f", Terms: ast.Terms{
					{Name: "p"},
					{Name: "n1", Value: "name"},
					{Name: "name"},
				}},
			},
		},
	}

	c := compiler.New()
	c.Compile(e)
	// reader := c.Reader("h")

	// for i := 0; i < bb.N; i++ {
	// 	for {
	// 		if err := reader.Read(nil); err != nil {
	// 			break
	// 		}
	// 	}
	// }
}
