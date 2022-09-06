package gen

import (
	"github.com/0xdbf/sigma/internal/vm"
)

/*

scalar ...
*/
type Scalar[A any] struct {
	addr vm.Addr
	seq  []A
	pos  int
}

/*


Scalar ...
*/
func NewScalar[A any](addr vm.Addr, xs []A) *Scalar[A] {
	return &Scalar[A]{addr: addr, seq: xs, pos: 0}
}

//
//
func (seq *Scalar[A]) Head(heap *vm.Heap) error {
	heap.Put(seq.addr, &seq.seq[seq.pos])
	return nil
}

//
func (seq *Scalar[A]) Tail(*vm.Heap) error {
	seq.pos++

	if len(seq.seq) == seq.pos {
		seq.pos = 0
		return vm.EOS
	}

	return nil
}
