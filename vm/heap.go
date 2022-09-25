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
