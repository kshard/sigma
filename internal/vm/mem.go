package vm

import "fmt"

/*

Addr ... (todo Addr)
*/
type Addr uint32

func (addr Addr) ReadOnly() Addr { return addr | (1 << 31) }

func (addr Addr) Writable() bool { return addr&(1<<31) == 0 }
func (addr Addr) Value() uint32  { return (uint32(addr) & 0x7fffffff) }
func (addr Addr) String() string {
	if addr.Writable() {
		return fmt.Sprintf("w%d", addr.Value())
	}

	return fmt.Sprintf("r%d", addr.Value())
}

/*

Heap ...
*/
type Heap []any

func (heap *Heap) Put(addr Addr, val any) {
	if !addr.Writable() {
		return
	}
	(*heap)[addr] = val
}

func (heap *Heap) Get(addr Addr) any {
	if addr.Writable() {
		return nil
	}

	return (*heap)[addr.Value()]
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
