package horn

import (
	"fmt"
)

var EOS = fmt.Errorf("end of stream")

type Atom string

type Heap map[Atom]any

func (heap *Heap) Join(r Atom, x any) {
	(*heap)[r] = x
	// *heap = Heap(int(*heap) + x)
}

type Stream interface {
	Read(*Heap) error
	Skip()
}

// type Σ interface {
// 	Stream(Heap) Stream
// }

/*

horn ...
*/
type join struct {
	A, B Stream
	f    func(*Heap) error
}

func Join(A, B Stream) Stream {
	join := join{A: A, B: B}
	join.f = join.ReadAB
	return &join
}

func (join *join) ReadAB(heap *Heap) (err error) {
	err = join.A.Read(heap)
	if err != nil {
		return
	}

	join.B.Skip()
	err = join.B.Read(heap)
	if err != nil {
		return
	}

	join.f = join.ReadB
	return
}

func (join *join) ReadB(heap *Heap) (err error) {
	err = join.B.Read(heap)
	if err != nil {
		return join.ReadAB(heap)
	}

	return
}

func (join *join) Read(heap *Heap) (err error) {
	return join.f(heap)
}

func (join *join) Skip() {
	join.B.Skip()
	join.A.Skip()
	join.f = join.ReadAB
}

// func Horn(Stream) Stream {
// 	return &join{shape: shape, seq: seq}
// }

// func (h horn) Stream(heap Heap) Stream {
// 	return horn{seq: h.seq[1:]}.join(h.eval(heap))
// }

// func (h horn) eval(heap Heap) Stream {
// 	return Map(
// 		h.seq[0].Stream(heap),
// 		func(i Heap) Heap { return heap + i },
// 	)
// }

// func (h horn) join(sin Stream) Stream {
// 	if len(h.seq) == 0 {
// 		return sin
// 	}

// 	seg := FMap(sin, h.eval)
// 	return horn{seq: h.seq[1:]}.join(seg)
// }
