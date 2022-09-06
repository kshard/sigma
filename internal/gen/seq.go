package gen

import (
	"github.com/0xdbf/sigma/internal/vm"
)

/*

seq ...
*/
type Seq[A any] struct {
	addr []vm.Addr
	seq  [][]A
	pos  int
}

/*


Stream ...
*/
func NewSeq[A any](addr []vm.Addr, xs [][]A) *Seq[A] {
	return &Seq[A]{addr: addr, seq: xs, pos: 0}
}

//
//
func (seq *Seq[A]) Head(heap *vm.Heap) error {
	v := seq.seq[seq.pos]
	for i, addr := range seq.addr {
		heap.Put(addr, &v[i])
	}

	return nil
}

//
func (seq *Seq[A]) Tail(*vm.Heap) error {
	seq.pos++

	if len(seq.seq) == seq.pos {
		seq.pos = 0
		return vm.EOS
	}

	return nil
}
