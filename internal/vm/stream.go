package vm

import "errors"

var EOS = errors.New("end of stream")

/*

Stream ...

  for err := stream.Init(&h); err == nil; err = a.Read(&h) {
   ...
  }

*/
type Stream interface {
	// init stream & read head
	Init(*Heap) error
	// continue stream reading
	Read(*Heap) error
}

/*

streamFMap ...
*/
type fmap struct {
	head Stream
	tail Stream
}

/*

FMap ...
*/
func FMap(head, tail Stream) Stream {
	return &fmap{head: head, tail: tail}
}

func (fmap *fmap) Init(heap *Heap) error {
	if err := fmap.head.Init(heap); err != nil {
		return err
	}

	if err := fmap.tail.Init(heap); err != nil {
		return err
	}

	return nil
}

func (fmap *fmap) Read(heap *Heap) error {
	if err := fmap.tail.Read(heap); err == nil {
		return nil
	}

	if err := fmap.head.Read(heap); err != nil {
		return err
	}

	if err := fmap.tail.Init(heap); err != nil {
		return err
	}

	return nil
}
