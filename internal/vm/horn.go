package vm

func Horn(seq ...Stream) Stream {
	head := seq[0]
	for _, tail := range seq[1:] {
		head = FMap(head, tail)
	}
	return head
}

//
type eq struct {
	x      []Addr
	y      []Addr
	stream Stream
}

func Eq(x, y []Addr, stream Stream) Stream {
	return &eq{x: x, y: y, stream: stream}
}

// func (eq *eq) Head(heap *Heap) error {
// 	if err := eq.stream.Head(heap); err != nil {
// 		return err
// 	}

// 	for i, addr := range eq.x {
// 		vx := (*heap)[addr].(*any)
// 		vy := (*heap)[eq.y[i]].(*any)
// 		if *vx != *vy {
// 			return eq.Tail(heap)
// 		}

// 		// switch v := (*heap)[addr].(type) {
// 		// case *any:
// 		// 	if *v != eq.seq[i] {
// 		// 		return eq.Tail(heap)
// 		// 	}
// 		// }
// 	}

// 	return nil
// }

// func (eq *eq) Tail(heap *Heap) error {
// 	if err := eq.stream.Tail(heap); err != nil {
// 		return err
// 	}
// 	return eq.Head(heap)
// }

func (eq *eq) Init(heap *Heap) error {
	if err := eq.stream.Init(heap); err != nil {
		return err
	}

	for i, addr := range eq.x {
		vx := (*heap)[addr].(*any)
		vy := (*heap)[eq.y[i]].(*any)
		if *vx != *vy {
			return eq.Read(heap)
		}
	}

	return nil
}

func (eq *eq) Read(heap *Heap) error {
	if err := eq.stream.Read(heap); err != nil {
		return err
	}

	for i, addr := range eq.x {
		vx := (*heap)[addr].(*any)
		vy := (*heap)[eq.y[i]].(*any)
		if *vx != *vy {
			return eq.Read(heap)
		}
	}

	return nil
}
