package horn

// type Heap int

// type Stream interface {
// 	Head() Heap
// 	Tail() Stream
// }

// /*

// streamFMap ...
// */
// type streamFMap struct {
// 	stream Stream
// 	header Stream
// 	mapper func(Heap) Stream
// }

// /*

// FMap ...
// */
// func FMap(stream Stream, fmap func(Heap) Stream) Stream {
// 	return (streamFMap{stream: stream, header: nil, mapper: fmap}).Tail()
// }

// /*

// Head ...
// */
// func (fmap streamFMap) Head() Heap {
// 	return fmap.header.Head()
// }

// /*

// Tail ...
// TODO: recursive
// */
// func (fmap streamFMap) Tail() Stream {
// 	if fmap.header == nil && fmap.stream == nil {
// 		return nil
// 	}

// 	if fmap.header == nil {
// 		var head Stream
// 		tail := fmap.stream
// 		for head == nil {
// 			if tail == nil {
// 				return nil
// 			}
// 			head = fmap.mapper(tail.Head())
// 			tail = tail.Tail()
// 		}

// 		return streamFMap{stream: tail, header: head, mapper: fmap.mapper}
// 	}

// 	tail := fmap.header.Tail()
// 	if tail == nil {
// 		if fmap.stream == nil {
// 			return nil
// 		}

// 		return FMap(fmap.stream, fmap.mapper)
// 	}

// 	return streamFMap{stream: fmap.stream, header: tail, mapper: fmap.mapper}
// }

// /*

// map ...
// */
// type streamMap struct {
// 	stream Stream
// 	mapper func(Heap) Heap
// }

// /*

// Map ...
// */
// func Map(stream Stream, fmap func(Heap) Heap) Stream {
// 	if stream == nil {
// 		return nil
// 	}

// 	return streamMap{stream: stream, mapper: fmap}
// }

// /*

// Head ...
// */
// func (fmap streamMap) Head() Heap {
// 	return fmap.mapper(fmap.stream.Head())
// }

// /*

// Tail ...
// */
// func (fmap streamMap) Tail() Stream {
// 	tail := fmap.stream.Tail()
// 	if tail == nil {
// 		return nil
// 	}

// 	return streamMap{stream: tail, mapper: fmap.mapper}
// }

// /*

// ForEach ...
// */
// func ForEach(stream Stream) {
// 	if stream == nil {
// 		return
// 	}

// 	head := stream
// 	for {
// 		// fmt.Println(head.Head())
// 		head.Head()
// 		head = head.Tail()
// 		if head == nil {
// 			return
// 		}
// 	}
// }

func ForEach(stream Stream, heap *Heap) {
	stream.Skip()
	for {
		if err := stream.Read(heap); err != nil {
			return
		}
		// for k, v := range *heap {
		// 	fmt.Printf("%v: %v ", k, v)
		// }
		// fmt.Println()
		// fmt.Println(*heap)
	}

	// if stream == nil {
	// 	return
	// }

	// head := stream
	// for {
	// 	// fmt.Println(head.Head())
	// 	head.Head()
	// 	head = head.Tail()
	// 	if head == nil {
	// 		return
	// 	}
	// }
}
