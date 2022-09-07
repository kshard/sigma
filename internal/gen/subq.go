package gen

import (
	"github.com/0xdbf/sigma/internal/vm"
)

/*

seq ...
*/
type SubQ struct {
	addr []vm.Addr
	seq  [][]any
	pat  []any
	pos  int
}

/*


Stream ...
*/
func NewSubQ(addr []vm.Addr, xs [][]any) *SubQ {
	return &SubQ{
		addr: addr,
		seq:  xs,
		pat:  make([]any, len(addr)),
		pos:  0,
	}
}

func (seq *SubQ) Init(heap *vm.Heap) error {
	seq.pos = 0
	return seq.Read(heap)
}

func (seq *SubQ) Read(heap *vm.Heap) error {
	for i, addr := range seq.addr {
		if !addr.Writable() {
			seq.pat[i] = *heap.Get(addr).(*any)
		} else {
			seq.pat[i] = nil
		}
	}
	if err := seq.Skip(seq.pat); err != nil {
		return err
	}

	v := seq.seq[seq.pos]
	for i, addr := range seq.addr {
		if addr.Writable() {
			heap.Put(addr, &v[i])
		}
	}

	seq.pos++
	return nil
}

func (seq *SubQ) Skip(val []any) error {
	for {
		if len(seq.seq) == seq.pos {
			return vm.EOS
		}

		eq := true
		for i, x := range val {
			if x != nil && seq.seq[seq.pos][i] != x {
				eq = false
			}
		}

		if eq {
			return nil
		}

		seq.pos++
	}
}
