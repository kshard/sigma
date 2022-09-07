package vm_test

import (
	"testing"

	"github.com/0xdbf/sigma/internal/gen"
	"github.com/0xdbf/sigma/internal/vm"
)

func TestEq(t *testing.T) {
	director := vm.Eq(
		[]vm.Addr{1, 2}, []vm.Addr{3, 4},
		gen.NewSeq([]vm.Addr{0, 1, 2}, gen.IMDB()),
	)

	horn := vm.Horn(director)
	heap := make(vm.Heap, 5)
	heap[3] = ptr("name")
	heap[4] = ptr("Ridley Scott")

	vm.Debug(horn, &heap)
}

func ptr(x any) *any { return &x }

func TestXxx(t *testing.T) {
	/*
			Memory management (move? calls for f(...))

		   ?- h(_, _).

		   f(s, p, o).

		   h(s, title) :-
		      f(s, \"year\", 1987),     // s = ? p = year,  o = 1987
		      f(s, \"title\", title).   // s = $ p = title, o = ?
	*/

	year := vm.Eq(
		[]vm.Addr{1, 2}, []vm.Addr{3, 4},
		gen.NewSeq([]vm.Addr{0, 1, 2}, gen.IMDB()),
	)
	title := vm.Eq(
		[]vm.Addr{0, 6}, []vm.Addr{5, 8},
		gen.NewSeq([]vm.Addr{5, 6, 7}, gen.IMDB()),
	)

	horn := vm.Horn(year, title)
	heap := make(vm.Heap, 9)
	heap[3] = ptr("year")
	heap[4] = ptr(1987)
	heap[8] = ptr("title")

	vm.Debug(horn, &heap)

}

func TestTx(t *testing.T) {
	// movie := vm.Eq(
	// 	[]vm.Addr{1, 2}, []vm.Addr{3, 4},
	// 	gen.NewSeq([]vm.Addr{0, 1, 2}, gen.IMDB()),
	// )

	// cast := vm.Eq(
	// 	[]vm.Addr{0, 6}, []vm.Addr{5, 8},
	// 	gen.NewSeq([]vm.Addr{5, 6, 7}, gen.IMDB()),
	// )

	movie := gen.NewSubQ([]vm.Addr{0, 3 | (1 << 31), 4 | (1 << 31)}, gen.IMDB())
	cast := gen.NewSubQ([]vm.Addr{0 | (1 << 31), 8 | (1 << 31), 7}, gen.IMDB())
	name := gen.NewSubQ([]vm.Addr{7 | (1 << 31), 11 | (1 << 31), 10}, gen.IMDB())

	horn := vm.Horn(movie, cast, name)

	heap := make(vm.Heap, 13)
	heap[3] = ptr("title")
	heap[4] = ptr("Lethal Weapon")
	heap[8] = ptr("cast")
	heap[11] = ptr("name")

	vm.Debug(horn, &heap)
}

func BenchmarkTx(bb *testing.B) {
	movie := gen.NewSubQ([]vm.Addr{0, 3 | (1 << 31), 4 | (1 << 31)}, gen.IMDB())
	cast := gen.NewSubQ([]vm.Addr{0 | (1 << 31), 8 | (1 << 31), 7}, gen.IMDB())
	name := gen.NewSubQ([]vm.Addr{7 | (1 << 31), 11 | (1 << 31), 10}, gen.IMDB())

	horn := vm.Horn(movie, cast, name)

	heap := make(vm.Heap, 13)
	heap[3] = ptr("title")
	heap[4] = ptr("Lethal Weapon")
	heap[8] = ptr("cast")
	heap[11] = ptr("name")

	for i := 0; i < bb.N; i++ {
		vm.Eval(horn, &heap)
	}
}
