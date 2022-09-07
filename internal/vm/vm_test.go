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
	movie := vm.Eq(
		[]vm.Addr{1, 2}, []vm.Addr{3, 4},
		gen.NewSeq([]vm.Addr{0, 1, 2}, gen.IMDB()),
	)

	cast := vm.Eq(
		[]vm.Addr{0, 6}, []vm.Addr{5, 8},
		gen.NewSeq([]vm.Addr{5, 6, 7}, gen.IMDB()),
	)

	name := vm.Eq(
		[]vm.Addr{7, 10}, []vm.Addr{9, 12},
		gen.NewSeq([]vm.Addr{9, 10, 11}, gen.IMDB()),
	)

	horn := vm.Horn(movie, cast, name)

	heap := make(vm.Heap, 13)
	heap[3] = ptr("title")
	heap[4] = ptr("Lethal Weapon")
	heap[8] = ptr("cast")
	heap[12] = ptr("name")

	vm.Debug(horn, &heap)
}
