package gen

import (
	"github.com/0xdbf/sigma/internal/vm"
)

/*

seq ...
*/
type Seq struct {
	addr []vm.Addr
	seq  [][]any
	pos  int
}

/*


Stream ...
*/
func NewSeq(addr []vm.Addr, xs [][]any) *Seq {
	return &Seq{addr: addr, seq: xs, pos: 0}
}

func (seq *Seq) Init(heap *vm.Heap) error {
	seq.pos = 0
	return seq.Read(heap)
}

func (seq *Seq) Read(heap *vm.Heap) error {
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
