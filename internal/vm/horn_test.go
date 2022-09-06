package vm_test

import (
	"testing"

	"github.com/0xdbf/sigma/internal/gen"
	"github.com/0xdbf/sigma/internal/vm"
)

func TestJoin(t *testing.T) {
	a := gen.NewScalar(vm.Addr(0), []string{"|1|", "|2|", "|3|"})
	b := gen.NewScalar(vm.Addr(1), []int{4, 5, 6})
	c := gen.NewScalar(vm.Addr(2), []int{7, 8, 9})

	s := vm.Horn(a, b, c)
	h := make(vm.Heap, 3)

	vm.Debug(s, &h)
}

func BenchmarkJoinInt(bb *testing.B) {
	a := gen.NewScalar(vm.Addr(0), []int{1, 2, 3})
	b := gen.NewScalar(vm.Addr(1), []int{4, 5, 6})
	c := gen.NewScalar(vm.Addr(2), []int{7, 8, 9})

	s := vm.Horn(a, b, c)
	h := make(vm.Heap, 3)

	for i := 0; i < bb.N; i++ {
		vm.Eval(s, &h)
	}
}

func BenchmarkJoinString(bb *testing.B) {
	a := gen.NewScalar(vm.Addr(0), []string{"1", "2", "3"})
	b := gen.NewScalar(vm.Addr(1), []string{"4", "5", "6"})
	c := gen.NewScalar(vm.Addr(2), []string{"7", "8", "9"})

	s := vm.Horn(a, b, c)
	h := make(vm.Heap, 3)

	for i := 0; i < bb.N; i++ {
		vm.Eval(s, &h)
	}
}
