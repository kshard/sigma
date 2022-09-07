package vm_test

import (
	"testing"

	"github.com/0xdbf/sigma/internal/gen"
	"github.com/0xdbf/sigma/internal/vm"
)

func TestJoin(t *testing.T) {
	a := gen.NewScalar(vm.Addr(0), []any{"|1|", "|2|", "|3|"})
	b := gen.NewScalar(vm.Addr(1), []any{4, 5, 6})
	c := gen.NewScalar(vm.Addr(2), []any{7, 8, 9})

	s := vm.Horn(a, b, c)
	h := make(vm.Heap, 3)

	for ea := s.Init(&h); ea == nil; ea = s.Read(&h) {
		h.Dump()
	}
	// vm.Debug(s, &h)
}

func TestAbc(t *testing.T) {
	a := gen.NewScalar(vm.Addr(0), []any{"|1|", "|2|", "|3|"})
	b := gen.NewScalar(vm.Addr(1), []any{4, 5, 6})
	c := gen.NewScalar(vm.Addr(2), []any{7, 8, 9})

	h := make(vm.Heap, 3)

	for ea := a.Init(&h); ea == nil; ea = a.Read(&h) {
		for eb := b.Init(&h); eb == nil; eb = b.Read(&h) {
			for ec := c.Init(&h); ec == nil; ec = c.Read(&h) {
				h.Dump()
			}
		}
	}

	// How to make stream from this loops ?
	/*
		a.Read(&h)
		b.Read(&h)
		c.Read(&h)
		...
	*/

}

func BenchmarkNestedLoop(bb *testing.B) {
	a := gen.NewScalar(vm.Addr(0), []any{1, 2, 3})
	b := gen.NewScalar(vm.Addr(1), []any{4, 5, 6})
	c := gen.NewScalar(vm.Addr(2), []any{7, 8, 9})

	h := make(vm.Heap, 3)

	for i := 0; i < bb.N; i++ {
		for ea := a.Init(&h); ea == nil; ea = a.Read(&h) {
			for eb := b.Init(&h); eb == nil; eb = b.Read(&h) {
				for ec := c.Init(&h); ec == nil; ec = c.Read(&h) {

				}
			}
		}
	}

}

func BenchmarkJoinInt(bb *testing.B) {
	a := gen.NewScalar(vm.Addr(0), []any{1, 2, 3})
	b := gen.NewScalar(vm.Addr(1), []any{4, 5, 6})
	c := gen.NewScalar(vm.Addr(2), []any{7, 8, 9})

	s := vm.Horn(a, b, c)
	h := make(vm.Heap, 3)

	for i := 0; i < bb.N; i++ {
		for ea := s.Init(&h); ea == nil; ea = s.Read(&h) {
		}
		// vm.Eval(s, &h)
	}
}

func BenchmarkJoinString(bb *testing.B) {
	a := gen.NewScalar(vm.Addr(0), []any{"1", "2", "3"})
	b := gen.NewScalar(vm.Addr(1), []any{"4", "5", "6"})
	c := gen.NewScalar(vm.Addr(2), []any{"7", "8", "9"})

	s := vm.Horn(a, b, c)
	h := make(vm.Heap, 3)

	for i := 0; i < bb.N; i++ {
		vm.Eval(s, &h)
	}
}
