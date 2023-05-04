/*

  Sigma Virtual Machine
  Copyright (C) 2016 - 2023 Dmitry Kolesnikov

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

import "github.com/kshard/xsd"

//
// The file defines stream reader interface.
//

// VM returns Reader as a result of σ-application.
// VM associates reader object with goal of the application (σ-stream) that
// allows a client to use `interface{ Read([]any) error }` as lazy consumer of
// the relation.
type Reader struct {
	stream Stream
	heap   *Heap
	addr   []Addr
	closed bool
}

// ToSeq evaluates the stream and copy context into sequence of relation
func (reader *Reader) ToSeq() [][]xsd.Value {
	seq := [][]xsd.Value{}

	for {
		val := make([]xsd.Value, len(reader.addr))
		if err := reader.Read(val); err != nil {
			break
		}
		seq = append(seq, val)
	}

	return seq
}

// Read "car" (head of the stream) into container
func (reader *Reader) Read(seq []xsd.Value) error {
	if reader.closed {
		if err := reader.stream.Init(reader.heap); err != nil {
			return err
		}
		reader.closed = false
		reader.copyHead(seq)
		return nil
	}

	if err := reader.stream.Read(reader.heap); err != nil {
		reader.closed = true
		return err
	}
	reader.copyHead(seq)
	return nil
}

// VM's heap is only the snapshot of current state.
// The heap values MUST be copied on the return to client.
func (reader *Reader) copyHead(seq []xsd.Value) {
	if seq == nil {
		return
	}

	for i, addr := range reader.addr {
		seq[i] = reader.heap.Get(addr)
		// switch x := reader.heap.Get(addr).(type) {
		// case *any:
		// 	seq[i] = *x
		// default:
		// 	seq[i] = x
		// }
	}
}
