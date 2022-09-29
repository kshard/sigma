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

package vm

import "fmt"

//
// The file defines random access memory (heap) model for VM.
//

/*

Heap of VM
*/
type Heap []any

// Put writes value to heap by the address
func (heap *Heap) Put(addr Addr, val any) {
	if !addr.IsWritable() {
		return
	}
	(*heap)[addr] = val
}

// Get reads value from heap
func (heap *Heap) Get(addr Addr) any {
	if addr.IsWritable() {
		return nil
	}

	return (*heap)[addr.Value()]
}

func (heap *Heap) UnsafeGet(addr Addr) any {
	return (*heap)[addr.Value()]
}

// Dump heaps on console
func (heap *Heap) Dump() {
	fmt.Print("[")
	for _, v := range *heap {
		switch x := v.(type) {
		case *any:
			fmt.Printf(" %v ", *x)
		default:
			fmt.Printf(" %v ", x)
		}
	}
	fmt.Println("]")
}
