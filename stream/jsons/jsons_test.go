package jsons_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/0xdbf/sigma"
	"github.com/0xdbf/sigma/core"
	"github.com/0xdbf/sigma/stream"
	"github.com/0xdbf/sigma/stream/jsons"
)

// [1, 2, 3]
func TestHorn(t *testing.T) {
	val := `
	{
		"a": {"b": 1}
	}
	`
	var gen jsons.Value
	json.Unmarshal([]byte(val), &gen)

	seq1 := sigma.Head(
		[]core.Value{core.Atom("a"), core.Atom("b"), core.Atom("c")},
		jsons.Σ(),
	)
	seq2 := sigma.Head(
		[]core.Value{core.Atom("c"), core.Atom("d"), core.Atom("e")},
		jsons.Σ(),
	)
	seq := sigma.Horn([]core.Value{core.Atom("e")}, seq1, seq2)

	out := seq.Stream(core.Map{
		core.Atom("a"): gen,
		core.Atom("b"): core.Atom("a"),
		core.Atom("d"): core.Atom("b"),
	})

	stream.ForEach(
		stream.Map(out, func(v core.Values) core.Values {
			fmt.Println("~~> ", v)
			return nil
		}),
	)
}

func BenchmarkHorn(b *testing.B) {
	val := `
	{
		"a": {"b": 1}
	}
	`
	var gen jsons.Value
	json.Unmarshal([]byte(val), &gen)

	seq1 := sigma.Head(
		[]core.Value{core.Atom("a"), core.Atom("b"), core.Atom("c")},
		jsons.Σ(),
	)
	seq2 := sigma.Head(
		[]core.Value{core.Atom("c"), core.Atom("d"), core.Atom("e")},
		jsons.Σ(),
	)
	seq := sigma.Horn([]core.Value{core.Atom("e")}, seq1, seq2)

	out := seq.Stream(core.Map{
		core.Atom("a"): gen,
		core.Atom("b"): core.Atom("a"),
		core.Atom("d"): core.Atom("b"),
	})

	for i := 0; i < b.N; i++ {
		stream.ForEach(
			stream.Map(out, func(v core.Values) core.Values {
				return nil
			}),
		)
	}
}
