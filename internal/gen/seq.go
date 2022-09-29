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
	"github.com/0xdbf/sigma/vm"
)

// seq type
type Seq struct {
	addr []vm.Addr
	seq  [][]any
	pos  int
}

/*

Seq generates sequence of values
*/
func NewSeq(xs [][]any) func(...vm.Addr) *Seq {
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
		return vm.EOS
	}

	v := seq.seq[seq.pos]
	for i, addr := range seq.addr {
		heap.Put(addr, &v[i])
	}

	seq.pos++
	return nil
}

//
type Values struct {
	addr vm.Addr
	seq  []any
	pos  int
}

/*

Values generates sequence of values
*/
func NewValues(xs []any) func(...vm.Addr) *Values {
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
		return vm.EOS
	}

	heap.Put(seq.addr, &seq.seq[seq.pos])

	seq.pos++
	return nil
}
