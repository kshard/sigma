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
