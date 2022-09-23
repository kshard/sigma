package vm_test

import (
	"reflect"
	"testing"

	"github.com/0xdbf/sigma/internal/gen"
	"github.com/0xdbf/sigma/internal/vm"
)

func TestVmJoin(t *testing.T) {
	e := gen.NewSeq([][]any{})
	a := gen.NewSeq([][]any{{1}, {2}, {3}})
	b := gen.NewSeq([][]any{{4}, {5}, {6}})

	t.Run("TwoSeq", func(t *testing.T) {
		expect := [][]any{
			{1, 4}, {1, 5}, {1, 6},
			{2, 4}, {2, 5}, {2, 6},
			{3, 4}, {3, 5}, {3, 6},
		}

		seq := vm.New(2).
			Run([]vm.Addr{0, 1},
				vm.Horn(a(vm.Addr(0)), b(vm.Addr(1))),
			)

		if !reflect.DeepEqual(seq, expect) {
			t.Errorf("unexpected join of stream")
		}
	})

	t.Run("TwoSeqLastEmpty", func(t *testing.T) {
		expect := [][]any{}

		seq := vm.New(2).Run([]vm.Addr{0, 1},
			vm.Horn(a(vm.Addr(0)), e(vm.Addr(1))),
		)

		if !reflect.DeepEqual(seq, expect) {
			t.Errorf("unexpected join of stream")
		}
	})

	t.Run("TwoSeqFirstEmpty", func(t *testing.T) {
		expect := [][]any{}

		seq := vm.New(2).Run([]vm.Addr{0, 1},
			vm.Horn(e(vm.Addr(0)), b(vm.Addr(1))),
		)

		if !reflect.DeepEqual(seq, expect) {
			t.Errorf("unexpected join of stream")
		}
	})

	t.Run("FewSeq", func(t *testing.T) {
		expect := [][]any{
			{1, 1, 1}, {1, 1, 2}, {1, 1, 3},
			{1, 2, 1}, {1, 2, 2}, {1, 2, 3},
			{1, 3, 1}, {1, 3, 2}, {1, 3, 3},
			{2, 1, 1}, {2, 1, 2}, {2, 1, 3},
			{2, 2, 1}, {2, 2, 2}, {2, 2, 3},
			{2, 3, 1}, {2, 3, 2}, {2, 3, 3},
			{3, 1, 1}, {3, 1, 2}, {3, 1, 3},
			{3, 2, 1}, {3, 2, 2}, {3, 2, 3},
			{3, 3, 1}, {3, 3, 2}, {3, 3, 3},
		}

		seq := vm.New(3).Run([]vm.Addr{0, 1, 2},
			vm.Horn(a(vm.Addr(0)), a(vm.Addr(1)), a(vm.Addr(2))),
		)

		if !reflect.DeepEqual(seq, expect) {
			t.Errorf("unexpected join of stream")
		}
	})
}

func BenchmarkVmJoin(bb *testing.B) {
	a := gen.NewValues([]any{1, 2, 3})
	b := gen.NewValues([]any{4, 5, 6})
	c := gen.NewValues([]any{7, 8, 9})

	bb.Run("NestedLoop", func(bb *testing.B) {
		sA := a(0)
		sB := b(1)
		sC := c(2)
		h := make(vm.Heap, 3)

		for i := 0; i < bb.N; i++ {
			for ea := sA.Init(&h); ea == nil; ea = sA.Read(&h) {
				for eb := sB.Init(&h); eb == nil; eb = sB.Read(&h) {
					for ec := sC.Init(&h); ec == nil; ec = sC.Read(&h) {

					}
				}
			}
		}
	})

	bb.Run("Seq", func(bb *testing.B) {
		s := vm.Horn(a(0), b(1), c(2))
		h := make(vm.Heap, 3)

		for i := 0; i < bb.N; i++ {
			for err := s.Init(&h); err == nil; err = s.Read(&h) {
			}
		}
	})
}

/*
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
	/
			Memory management (move? calls for f(...))

		   ?- h(_, _).

		   f(s, p, o).

		   h(s, title) :-
		      f(s, \"year\", 1987),     // s = ? p = year,  o = 1987
		      f(s, \"title\", title).   // s = $ p = title, o = ?
	/

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
		vm.Exec(horn, &heap)
	}
}
*/
