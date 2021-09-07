package sigma

import (
	"github.com/0xdbf/sigma/core"
	"github.com/0xdbf/sigma/stream"
)

//
type join struct {
	seq []Σ
}

// Join ...
func Join(seq ...Σ) Σ {
	return &join{seq: seq}
}

func (h join) Stream(heap core.Values) stream.Stream {
	seq := make([]stream.Stream, len(h.seq))

	for i, σ := range h.seq {
		seq[i] = σ.Stream(heap)
	}

	return stream.Join(seq)
}
