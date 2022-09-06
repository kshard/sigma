package sigma

import (
	"github.com/0xdbf/sigma/core"
	"github.com/0xdbf/sigma/stream"
)

/*

Σ .... (sigma function implemented by data source / generator)
*/
type Σ interface {
	Stream(q core.Values) stream.Stream
}

/*

person(foaf:name, foaf:title) => Σ

professor(name) :- person(name, “prof”).

university(employee) :- professor(employee).


person(foaf:name, foaf:title).
--
memory frame [foaf:name, foaf:title = "prof"]
|
stream --> memory frame [foaf:name, foaf:title]

professor(name) :- person(name, “prof”).
--
memory [name]


type Stream interface {
	Read(core.Memory) error
}


ast.Generator{
	Name: ast.Atom("person"),
	Head: &ast.Terms{
		ast.Term("foaf:name"),
		ast.Term("foaf:title"),
	}
}

ast.Horn{
	Name: ast.Atom("professor"),
	Head: ast.Terms{
		ast.Term("name")
	},
	Body: ast.Body{
		ast.Predicate{
			Name: ast.Atom("person"),
			XXXX: ast.XXXX{
				ast.Term("name"),
				ast.Lit("prof")
			}
		}
	}
}

sigma.Compile(p).Eval(memory)

https://www.slideshare.net/feyeleanor/implementing-virtual-machines-in-go-c


*/
