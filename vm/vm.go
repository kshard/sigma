package vm

// VM instance
type VM struct {
	Heap Heap
}

// New creates instance of VM
func New(memSize int) *VM {
	return &VM{
		Heap: make(Heap, memSize),
	}
}

//
func (vm *VM) Stream(head []Addr, stream Stream) *Reader {
	addr := make([]Addr, len(head))
	for i, x := range head {
		addr[i] = x.ReadOnly()
	}

	return &Reader{
		stream: stream,
		heap:   &vm.Heap,
		addr:   addr,
		closed: true,
	}
}

// Run the logical program
func (vm *VM) Run(head []Addr, stream Stream) [][]any {
	seq := [][]any{}

	sio := vm.Stream(head, stream)
	for {
		val := make([]any, len(head))
		if err := sio.Read(val); err != nil {
			break
		}
		seq = append(seq, val)
	}

	return seq
}

func Horn(seq ...Stream) Stream {
	head := seq[0]
	for _, tail := range seq[1:] {
		head = Join(head, tail)
	}
	return head
}

//
//
type Reader struct {
	stream Stream
	heap   *Heap
	addr   []Addr
	closed bool
}

func (reader *Reader) ToSeq() [][]any {
	seq := [][]any{}

	for {
		val := make([]any, len(reader.addr))
		if err := reader.Read(val); err != nil {
			break
		}
		seq = append(seq, val)
	}

	return seq
}

func (reader *Reader) Read(seq []any) error {
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

func (reader *Reader) copyHead(seq []any) {
	if seq == nil {
		return
	}

	for i, addr := range reader.addr {
		switch x := reader.heap.Get(addr).(type) {
		case *any:
			seq[i] = *x
		default:
			seq[i] = x
		}
	}
}
