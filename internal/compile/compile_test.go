package compile_test

import (
	"fmt"
	"testing"

	"github.com/0xdbf/sigma/ast"
	"github.com/0xdbf/sigma/internal/compile"
	"github.com/0xdbf/sigma/internal/gen"
	"github.com/0xdbf/sigma/internal/vm"
)

func TestX(t *testing.T) {
	e := ast.Rules{
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

	c := compile.New()
	c.Facts["f"] = func(addr []vm.Addr) vm.Stream {
		return gen.NewSubQ(addr, gen.IMDB())
	}

	c.Compile(e)
	value := make([]any, 2)
	reader := c.Create([]string{"name", "name1"}, "h")
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

	c := compile.New()
	c.Facts["f"] = func(addr []vm.Addr) vm.Stream {
		return gen.NewSubQ(addr, gen.IMDB())
	}

	c.Compile(e)
	reader := c.Create([]string{"name"}, "h")

	for i := 0; i < bb.N; i++ {
		for {
			if err := reader.Read(nil); err != nil {
				break
			}
		}
	}
}
