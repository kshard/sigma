package horn

import (
	"testing"
)

func TestXxx(t *testing.T) {
	a := IStream(Atom("x"), []int{1, 2, 3})
	b := IStream(Atom("y"), []int{4, 5, 6})
	c := IStream(Atom("z"), []int{7, 8, 9})

	j := Join(a, Join(b, c))
	h := Heap(map[Atom]any{})

	ForEach(j, &h)
}

func BenchmarkXxx(bb *testing.B) {
	a := IStream(Atom("x"), []int{1, 2, 3})
	b := IStream(Atom("y"), []int{4, 5, 6})
	c := IStream(Atom("z"), []int{7, 8, 9})

	j := Join(a, Join(b, c))
	h := Heap(map[Atom]any{})

	for i := 0; i < bb.N; i++ {
		ForEach(j, &h)
	}
}
