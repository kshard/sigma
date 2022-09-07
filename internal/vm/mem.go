package vm

import "fmt"

/*

Addr ... (todo Addr)
*/
type Addr uint32

/*

Heap ...
*/
type Heap []any

func (heap *Heap) Put(addr Addr, val any) {
	if addr>>31 == 1 {
		return
	}
	(*heap)[addr] = val
}

func (heap *Heap) Get(addr Addr) any {
	return (*heap)[addr&0x7fffffff]
}

func (heap *Heap) Dump() {
	fmt.Print("[")
	for _, v := range *heap {
		switch x := v.(type) {
		case *string:
			fmt.Printf(" %v ", *x)
		case *int:
			fmt.Printf(" %v ", *x)
		case *any:
			fmt.Printf(" %v ", *x)
		default:
			fmt.Printf(" %v ", x)
		}
	}
	fmt.Println("]")
}

/*

?- movie(_, _).

// generator
imdb:movie(\"rdf:id\", \"schema:title\", \"schema:year\").

movie(id, title) :-
  imdb:movie(id, title, 1987).


?- h(_, _).

f(s, p, o).

h(s, o) :-
  f(s, \"name\", o), o = \"Ridley Scott\".


actors(id, name).
movies(id, title, year, cast).

casting(title, name) :-
  movies(_, title, year, cast),
	actors(cast, name),
	year < 1984.

*/
