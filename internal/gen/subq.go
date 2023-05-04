/*

  Sigma Virtual Machine
  Copyright (C) 2016  Dmitry Kolesnikov

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as published
  by the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/

package gen

import (
	"github.com/kshard/sigma/vm"
	"github.com/kshard/xsd"
)

/*
seq ...
*/
type SubQ struct {
	addr []vm.Addr
	seq  [][]xsd.Value
	pat  []xsd.Value
	pos  int
}

/*
Stream ...
*/
func NewSubQ(addr []vm.Addr, xs [][]xsd.Value) *SubQ {
	return &SubQ{
		addr: addr,
		seq:  xs,
		pat:  make([]xsd.Value, len(addr)),
		pos:  0,
	}
}

func (seq *SubQ) Init(heap *vm.Heap) error {
	// build sub-query
	for i, addr := range seq.addr {
		if !addr.IsWritable() {
			seq.pat[i] = heap.Get(addr)
		} else {
			seq.pat[i] = nil
		}
	}

	seq.pos = 0
	return seq.Read(heap)
}

func (seq *SubQ) Read(heap *vm.Heap) error {
	if err := seq.Skip(seq.pat); err != nil {
		return err
	}

	v := seq.seq[seq.pos]
	for i, addr := range seq.addr {
		if addr.IsWritable() {
			heap.Put(addr, v[i])
		}
	}

	seq.pos++
	return nil
}

func (seq *SubQ) Skip(val []xsd.Value) error {
	for {
		if len(seq.seq) == seq.pos {
			return vm.EndOfStream
		}

		eq := true
		for i, x := range val {
			if x != nil && xsd.Compare(seq.seq[seq.pos][i], x) != 0 {
				eq = false
			}
		}

		if eq {
			return nil
		}

		seq.pos++
	}
}
