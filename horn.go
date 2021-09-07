package sigma

import (
	"fmt"

	"github.com/0xdbf/sigma/core"
	"github.com/0xdbf/sigma/stream"
)

/*

horn ...
*/
type horn struct {
	shape core.Seq
	seq   []Σ
}

func Horn(shape core.Seq, seq ...Σ) Σ {
	return &horn{shape: shape, seq: seq}
}

func (h horn) Stream(heap core.Values) stream.Stream {
	switch heapV := heap.(type) {
	case core.Map:
		return stream.Map(
			horn{seq: h.seq[1:]}.join(h.eval(heapV)),
			func(v core.Values) core.Values {
				seq := make(core.Seq, len(h.shape))
				seq.Copy(h.shape, v)
				return seq
			},
		)
	}

	panic(fmt.Errorf("Unsupported memory type"))
}

func (h horn) eval(heap core.Map) stream.Stream {
	return stream.Map(h.seq[0].Stream(heap),
		func(v core.Values) core.Values {
			return v.Merge(heap)
		},
	)
}

func (h horn) join(sin stream.Stream) stream.Stream {
	if len(h.seq) == 0 {
		return sin
	}

	seg := stream.FMap(sin,
		func(x core.Values) stream.Stream {
			switch v := x.(type) {
			case core.Map:
				return h.eval(v)
			}
			panic(fmt.Errorf("Unsupported memory type"))
		},
	)

	return horn{seq: h.seq[1:]}.join(seg)
}
