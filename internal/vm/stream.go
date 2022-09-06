package vm

import "errors"

var EOS = errors.New("end of stream")

type Stream interface {
	Head(*Heap) error
	Tail(*Heap) error
}

/*

streamFMap ...
*/
type fmap struct {
	head   Stream
	tail   Stream
	closed bool
}

/*

FMap ...
*/
func FMap(head, tail Stream) Stream {
	return &fmap{head: head, tail: tail, closed: true}
}

/*

Head ...
*/
func (fmap *fmap) Head(heap *Heap) error {
	if fmap.closed {
		if err := fmap.head.Head(heap); err != nil {
			return err
		}
		fmap.closed = false
	}

	return fmap.tail.Head(heap)
}

/*

Tail ...
*/
func (fmap *fmap) Tail(heap *Heap) error {
	if err := fmap.tail.Tail(heap); err == nil {
		return nil
	}

	if err := fmap.head.Tail(heap); err != nil {
		return err
	}

	if err := fmap.head.Head(heap); err != nil {
		return err
	}

	return nil
}
