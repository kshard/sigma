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

func TestHorn(t *testing.T) {
	val := `
	{
		"a": {"b": [1, 2, 3]}
	}
	`
	var gen jsons.Value
	json.Unmarshal([]byte(val), &gen)

	x := core.Atom("x")
	y := core.Atom("y")
	z := core.Atom("z")

	seq := sigma.Head([]core.Value{x, y, z}, jsons.Σ())

	out := seq.Stream(core.Map{
		core.Atom("x"): gen,
		core.Atom("y"): core.Atom("a"),
		core.Atom("z"): core.Any,
	})

	stream.ForEach(
		stream.Map(out, func(v core.Values) core.Values {
			fmt.Println(v)
			return nil
		}),
	)
}
