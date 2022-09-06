package stream

import (
	"github.com/0xdbf/sigma/core"
)

/*

Stream ...
*/
type Stream interface {
	Head() core.Values
	Tail() Stream
}

/*

 */
type Unit core.Seq

func (unit Unit) Head() core.Values {
	return core.Seq(unit)
}

func (unit Unit) Tail() Stream {
	return nil
}

/*

 */
type streamFilter struct {
	stream Stream
	filter func(core.Values) bool
}

/*

Filter ...
*/
func Filter(stream Stream, filter func(core.Values) bool) Stream {
	tail := stream

	for !filter(tail.Head()) {
		tail = tail.Tail()
		if tail == nil {
			return nil
		}
	}

	return &streamFilter{stream: tail, filter: filter}
}

func (op *streamFilter) Head() core.Values {
	return op.stream.Head()
}

func (op *streamFilter) Tail() Stream {
	tail := op.stream.Tail()
	if tail == nil {
		return nil
	}

	for !op.filter(tail.Head()) {
		tail = tail.Tail()
		if tail == nil {
			return nil
		}
	}

	return &streamFilter{stream: tail, filter: op.filter}
}

/*

map ...
*/
type streamMap struct {
	stream Stream
	mapper func(core.Values) core.Values
}

/*

Map ...
*/
func Map(stream Stream, fmap func(core.Values) core.Values) Stream {
	if stream == nil {
		return nil
	}

	return &streamMap{stream: stream, mapper: fmap}
}

/*

Head ...
*/
func (fmap *streamMap) Head() core.Values {
	return fmap.mapper(fmap.stream.Head())
}

/*

Tail ...
*/
func (fmap *streamMap) Tail() Stream {
	tail := fmap.stream.Tail()
	if tail == nil {
		return nil
	}

	return &streamMap{stream: tail, mapper: fmap.mapper}
}

/*

streamFMap ...
*/
type streamFMap struct {
	stream Stream
	header Stream
	mapper func(core.Values) Stream
}

/*

FMap ...
*/
func FMap(stream Stream, fmap func(core.Values) Stream) Stream {
	return (&streamFMap{stream: stream, header: nil, mapper: fmap}).Tail()
}

/*

Head ...
*/
func (fmap *streamFMap) Head() core.Values {
	return fmap.header.Head()
}

/*

Tail ...
TODO: recursive
*/
func (fmap *streamFMap) Tail() Stream {
	if fmap.header == nil && fmap.stream == nil {
		return nil
	}

	if fmap.header == nil {
		var head Stream
		tail := fmap.stream
		for head == nil {
			if tail == nil {
				return nil
			}
			head = fmap.mapper(tail.Head())
			tail = tail.Tail()
		}

		return &streamFMap{stream: tail, header: head, mapper: fmap.mapper}
	}

	tail := fmap.header.Tail()
	if tail == nil {
		if fmap.stream == nil {
			return nil
		}

		return FMap(fmap.stream, fmap.mapper)
	}

	return &streamFMap{stream: fmap.stream, header: tail, mapper: fmap.mapper}
}

type streamJoin struct {
	seq []Stream
}

func Join(seq []Stream) Stream {
	return &streamJoin{seq: seq}
}

/*

Head ...
*/
func (join *streamJoin) Head() core.Values {
	return join.seq[0].Head()
}

func (join *streamJoin) Tail() Stream {
	if len(join.seq) == 0 {
		return nil
	}

	tail := join.seq[0].Tail()
	if tail == nil {
		if len(join.seq) == 1 {
			return nil
		}
		return Join(join.seq[1:])
	}

	seq := []Stream{tail}
	seq = append(seq, join.seq[1:]...)

	return Join(seq)
}

/*

ForEach ...
*/
func ForEach(stream Stream) {
	if stream == nil {
		return
	}

	head := stream
	for {
		head.Head()
		head = head.Tail()
		if head == nil {
			return
		}
	}
}
