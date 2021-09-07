package ints

import (
	"github.com/0xdbf/sigma"
	"github.com/0xdbf/sigma/core"
	"github.com/0xdbf/sigma/stream"
)

/*

ints ...
*/
type ints struct {
	seq [][]int
}

/*


Stream ...
*/
func Stream(seq [][]int) stream.Stream {
	return ints{seq: seq}
}

func (seq ints) Head() core.Values {
	values := make(core.Seq, len(seq.seq[0]))
	for k, v := range seq.seq[0] {
		values[k] = core.Int(v)
	}
	return values
}

func (seq ints) Tail() stream.Stream {
	if len(seq.seq) == 1 {
		return nil
	}

	return ints{seq: seq.seq[1:]}
}

/*


Σ ...
*/
type σ struct {
	seq [][]int
}

/*

Σ ...
*/
func Σ(seq [][]int) sigma.Σ {
	return σ{seq: seq}
}

func (σ σ) Stream(q core.Values) stream.Stream {
	sio := Stream(σ.seq)

	for i := 0; i < q.Len(); i++ {
		pat := q.Get(core.Int(i))
		if !pat.Eq(core.Any) {
			ind := i
			sio = stream.Filter(sio,
				func(v core.Values) bool {
					return v.Get(core.Int(ind)).Eq(pat)
				})
		}
	}

	return sio
}
