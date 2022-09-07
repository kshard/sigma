package gen

import (
	"github.com/0xdbf/sigma/internal/vm"
)

/*

scalar ...
*/
type Scalar struct {
	addr vm.Addr
	seq  []any
	pos  int
}

/*


Scalar ...
*/
func NewScalar(addr vm.Addr, xs []any) *Scalar {
	return &Scalar{addr: addr, seq: xs, pos: 0}
}

//
//
func (seq *Scalar) Head(heap *vm.Heap) error {
	heap.Put(seq.addr, &seq.seq[seq.pos])
	return nil
}

//
func (seq *Scalar) Tail(*vm.Heap) error {
	seq.pos++

	if len(seq.seq) == seq.pos {
		seq.pos = 0
		return vm.EOS
	}

	return nil
}

func (seq *Scalar) Init(heap *vm.Heap) error {
	seq.pos = 0
	return seq.Read(heap)
}

func (seq *Scalar) Read(heap *vm.Heap) error {
	if len(seq.seq) == seq.pos {
		return vm.EOS
	}

	heap.Put(seq.addr, &seq.seq[seq.pos])

	seq.pos++
	return nil
}
