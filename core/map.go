package core

/*

Map ...
*/
type Map map[Value]Value

/*

Merge ...
*/
func (heap Map) Merge(values Values) Values {
	switch v := values.(type) {
	case Map:
		return heap.merge(v)
	}

	return heap
}

func (heap Map) merge(x Map) Values {
	for key, val := range x {
		_, exists := heap[key]
		if !exists {
			heap[key] = val
		}
	}
	return heap
}

/*

Len ...
*/
func (heap Map) Len() int {
	return len(heap)
}

/*

Get ...
*/
func (heap Map) Get(key Value) Value {
	val, exists := heap[key]
	if !exists {
		return Any
	}

	return val
}

/*

Copy ...
*/
func (heap Map) Copy(keys []Value, values Values) Values {
	for i, key := range keys {
		heap[key] = values.Get(Int(i))
	}
	return heap
}
