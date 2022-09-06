package sigma

import (
	"fmt"

	"github.com/0xdbf/sigma/core"
	"github.com/0xdbf/sigma/stream"
)

//
//
//
type head struct {
	shape []core.Value
	sigma Σ
}

func Head(shape []core.Value, sigma Σ) Σ {
	return head{shape: shape, sigma: sigma}
}

func (h head) Stream(heap core.Values) stream.Stream {
	switch heapV := heap.(type) {
	case core.Map:
		return h.stream(heapV)
	}

	panic(fmt.Errorf("Unsupported memory type"))
}

func (h head) stream(heap core.Map) stream.Stream {
	seq := make(core.Seq, len(h.shape))
	seq.Copy(h.shape, heap)

	return stream.Map(h.sigma.Stream(seq),
		func(v core.Values) core.Values {
			return core.Map{}.Copy(h.shape, v)
		},
	)
}
