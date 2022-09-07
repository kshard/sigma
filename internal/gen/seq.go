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

func (seq *Seq[A]) Init(heap *vm.Heap) error {
	seq.pos = 0
	return seq.Read(heap)
}

func (seq *Seq[A]) Read(heap *vm.Heap) error {
	if len(seq.seq) == seq.pos {
		return vm.EOS
	}

	v := seq.seq[seq.pos]
	for i, addr := range seq.addr {
		heap.Put(addr, &v[i])
	}

	seq.pos++
	return nil
}
