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

// seq type
type Seq struct {
	addr []vm.Addr
	seq  [][]xsd.Value
	pos  int
}

/*
Seq generates sequence of values
*/
func NewSeq(xs [][]xsd.Value) func(...vm.Addr) *Seq {
	return func(addr ...vm.Addr) *Seq {
		return &Seq{addr: addr, seq: xs, pos: 0}
	}
}

func (seq *Seq) Init(heap *vm.Heap) error {
	seq.pos = 0
	return seq.Read(heap)
}

func (seq *Seq) Read(heap *vm.Heap) error {
	if len(seq.seq) == seq.pos {
		return vm.EndOfStream
	}

	v := seq.seq[seq.pos]
	for i, addr := range seq.addr {
		heap.Put(addr, v[i])
	}

	seq.pos++
	return nil
}

type Values struct {
	addr vm.Addr
	seq  []xsd.Value
	pos  int
}

/*
Values generates sequence of values
*/
func NewValues(xs []xsd.Value) func(...vm.Addr) *Values {
	return func(addr ...vm.Addr) *Values {
		return &Values{addr: addr[0], seq: xs, pos: 0}
	}
}

func (seq *Values) Init(heap *vm.Heap) error {
	seq.pos = 0
	return seq.Read(heap)
}

func (seq *Values) Read(heap *vm.Heap) error {
	if len(seq.seq) == seq.pos {
		return vm.EndOfStream
	}

	heap.Put(seq.addr, seq.seq[seq.pos])

	seq.pos++
	return nil
}
