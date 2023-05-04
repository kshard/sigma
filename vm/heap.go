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
	"fmt"

	"github.com/kshard/xsd"
)

//
// The file defines random access memory (heap) model for VM.
//

// Heap of VM
type Heap []xsd.Value

// Put writes value to heap by the address
func (heap *Heap) Put(addr Addr, val xsd.Value) {
	if !addr.IsWritable() {
		return
	}
	(*heap)[addr] = val
}

// Get reads value from heap
func (heap *Heap) Get(addr Addr) xsd.Value {
	if addr.IsWritable() {
		return nil
	}

	return (*heap)[addr.Value()]
}

func (heap *Heap) UnsafeGet(addr Addr) xsd.Value {
	return (*heap)[addr.Value()]
}

// Dump heaps on console
func (heap *Heap) Dump() {
	fmt.Print("[")
	for _, v := range *heap {
		fmt.Printf(" %v ", v)
		// switch x := v.(type) {
		// case *any:
		// 	fmt.Printf(" %v ", *x)
		// default:
		// 	fmt.Printf(" %v ", x)
		// }
	}
	fmt.Println("]")
}
