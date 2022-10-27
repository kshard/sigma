/*

  Copyright 2016 Dmitry Kolesnikov, All Rights Reserved

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package gen

import (
	"github.com/kshard/sigma/vm"
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
	// build sub-query
	for i, addr := range seq.addr {
		if !addr.IsWritable() {
			seq.pat[i] = *heap.Get(addr).(*any)
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
			heap.Put(addr, &v[i])
		}
	}

	seq.pos++
	return nil
}

func (seq *SubQ) Skip(val []any) error {
	for {
		if len(seq.seq) == seq.pos {
			return vm.ErrEndOfStream
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
