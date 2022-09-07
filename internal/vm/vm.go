package vm

func Eval(stream Stream, heap *Heap) (err error) {
	for err = stream.Init(heap); err == nil; err = stream.Read(heap) {
	}

	return
}

func Debug(stream Stream, heap *Heap) (err error) {
	for err = stream.Init(heap); err == nil; err = stream.Read(heap) {
		heap.Dump()
	}

	return
}
