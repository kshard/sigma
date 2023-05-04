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
	"fmt"
	"testing"

	"github.com/kshard/sigma/asm"
	"github.com/kshard/sigma/ast"
	"github.com/kshard/sigma/internal/compiler"
	"github.com/kshard/sigma/internal/gen"
	"github.com/kshard/xsd"
)

func TestX(t *testing.T) {
	rules := ast.Rules{
		ast.NewFact("f").Tuple("s", "p", "o"),

		ast.NewHorn(
			ast.NewHead("a").Tuple("movie", "cast"),
			ast.NewExpr("f").
				Term("m").
				Term("t1", xsd.From("title")).
				Term("movie"),
			ast.NewExpr("f").
				Term("m").
				Term("c1", xsd.From("cast")).
				Term("cast"),
		),

		ast.NewHorn(
			ast.NewHead("h").Tuple("name", "name1"),
			ast.NewExpr("a").
				Term("t2", xsd.From("Lethal Weapon")).
				Term("p"),
			ast.NewExpr("f").
				Term("p").
				Term("n1", xsd.From("name")).
				Term("name"),
			ast.NewExpr("a").
				Term("t3", xsd.From("Mad Max")).
				Term("s"),
			ast.NewExpr("f").
				Term("s").
				Term("n1", xsd.From("name")).
				Term("name1"),
		),
	}

	build := compiler.New()
	if err := build.Compile(rules); err != nil {
		panic(err)
	}

	machine, shape, code := build.Assemble("h")

	ctx := asm.NewContext().Add("f", gen.FactsIMDB)
	reader := machine.Stream(shape, code.Link(ctx))

	value := make([]xsd.Value, len(shape))

	for {
		if err := reader.Read(value); err != nil {
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
		ast.NewFact("f").Tuple("s", "p", "o"),

		&ast.Horn{
			Head: &ast.Head{Name: "h", Terms: ast.Terms{{Name: "name"}}},
			Body: []*ast.Imply{
				{Name: "f", Terms: ast.Terms{
					{Name: "m"},
					{Name: "t1", Value: xsd.From("title")},
					{Name: "t2", Value: xsd.From("Lethal Weapon")},
				}},
				{Name: "f", Terms: ast.Terms{
					{Name: "m"},
					{Name: "c1", Value: xsd.From("cast")},
					{Name: "p"},
				}},
				{Name: "f", Terms: ast.Terms{
					{Name: "p"},
					{Name: "n1", Value: xsd.From("name")},
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
