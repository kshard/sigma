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

package vm

import (
	"errors"

	"github.com/kshard/xsd"
)

//
// The file implements Sigma VM
//

// Generator is σ-expressions that define ground facts for
// the application context. Typically, Generator fetches ground facts from
// external storage.
type Generator func([]Addr) Stream

// Stream of relations is core abstraction used by the VM.
// Stream produces lazy sequence of relations (tuples).
//
//	for err := stream.Init(&h); err == nil; err = stream.Read(&h) {
//	 ...
//	}
type Stream interface {
	// init stream & read head
	Init(*Heap) error
	// continue stream reading
	Read(*Heap) error
}

// The constant value reports that End Of Stream
var EndOfStream = errors.New("end of stream")

// Sigma VM
type VM struct {
	Heap Heap
}

// New creates new instance of VM
func New(memSize int) *VM {
	return &VM{
		Heap: make(Heap, memSize),
	}
}

// Creates Reader for given Stream and evaluates its goals
func (vm *VM) Stream(head []Addr, stream Stream) *Reader {
	addr := make([]Addr, len(head))
	for i, x := range head {
		addr[i] = x.ReadOnly()
	}

	return &Reader{
		stream: stream,
		heap:   &vm.Heap,
		addr:   addr,
		closed: true,
	}
}

// Evaluates Stream, return sequence of relations
func (vm *VM) Run(head []Addr, stream Stream) [][]xsd.Value {
	return vm.Stream(head, stream).ToSeq()
}

// Horn clause corresponds to join (⨝) operator
func Horn(seq ...Stream) Stream {
	head := seq[0]
	for _, tail := range seq[1:] {
		head = Join(head, tail)
	}
	return head
}
