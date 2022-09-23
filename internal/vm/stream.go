package vm

//
// The file defines stream abstraction and its combinators.
//

import (
	"errors"
)

var EOS = errors.New("end of stream")

/*

Stream sequence of data elements made available over time.

  for err := stream.Init(&h); err == nil; err = stream.Read(&h) {
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

Join is a fundamental stream combinator, it builds a new stream by
evaluating a tail stream for each element of head stream.

  for err := head.Init(&h); err == nil; err = head.Read(&h) {
    for err := tail.Init(&h); err == nil; err = tail.Read(&h) {
      ...
    }
  }

*/
func Join(head, tail Stream) Stream {
	return &join{head: head, tail: tail}
}

/*

stream left join operator
*/
type join struct {
	head Stream
	tail Stream
}

func (fmap *join) Init(heap *Heap) error {
	if err := fmap.head.Init(heap); err != nil {
		return err
	}

	if err := fmap.tail.Init(heap); err != nil {
		return err
	}

	return nil
}

func (fmap *join) Read(heap *Heap) error {
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
