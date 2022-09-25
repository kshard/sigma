package vm

//
type eq struct {
	x      []Addr
	y      []Addr
	stream Stream
}

func Eq(x, y []Addr, stream Stream) Stream {
	return &eq{x: x, y: y, stream: stream}
}

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
