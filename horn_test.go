package sigma_test

import (
	"fmt"
	"testing"

	"github.com/0xdbf/sigma"
	"github.com/0xdbf/sigma/core"
	"github.com/0xdbf/sigma/stream"
	"github.com/0xdbf/sigma/stream/ints"
)

type T struct{ stream stream.Stream }

func (t T) Stream(q core.Seq) stream.Stream {
	fmt.Println("==> ", q)
	return t.stream
}

func TestHorn(t *testing.T) {
	x := core.Atom("x")
	y := core.Atom("y")
	z := core.Atom("z")
	b := core.Atom("b")

	seq := ints.Σ([][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
		{4, 5, 6},
		{5, 6, 7},
		{6, 7, 8},
		{7, 8, 9},
	})

	seq1 := sigma.Head([]core.Value{x, y}, seq)
	seq2 := sigma.Head([]core.Value{y, b, z}, seq)
	seq3 := sigma.Head([]core.Value{y, b, z}, seq)

	horn := sigma.Horn([]core.Value{x, y, z}, seq1, seq2, seq3)

	out := horn.Stream(core.Map{
		core.Atom("x"): core.Any, // stream.Int(3),
		core.Atom("y"): core.Any,
		core.Atom("z"): core.Any,
	})

	stream.ForEach(
		stream.Map(out, func(v core.Values) core.Values {
			fmt.Println(v)
			return nil
		}),
	)
}

// func BenchmarkHorn(bx *testing.B) {
// 	x := sigmavm.Atom("x")
// 	y := sigmavm.Atom("y")
// 	z := sigmavm.Atom("z")
// 	b := sigmavm.Atom("b")

// 	ints := stream.Ints([][]int{
// 		{1, 2, 3},
// 		{2, 3, 4},
// 		{3, 4, 5},
// 		{4, 5, 6},
// 		{5, 6, 7},
// 		{6, 7, 8},
// 		{7, 8, 9},
// 	})

// 	seq1 := sigmavm.Seq([]stream.Value{x, y}, ints)
// 	seq2 := sigmavm.Seq([]stream.Value{y, b, z}, ints)
// 	seq3 := sigmavm.Seq([]stream.Value{y, b, z}, ints)

// 	horn := sigmavm.Horn{seq1, seq2, seq3}

// 	out := horn.Stream(sigmavm.SubQuery{
// 		sigmavm.Atom("x"): sigmavm.Any, // stream.Int(3),
// 		sigmavm.Atom("y"): sigmavm.Any,
// 		sigmavm.Atom("z"): sigmavm.Any,
// 	})

// 	bx.Run("xxx", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			stream.ForEach(out)
// 		}
// 	})

// }
