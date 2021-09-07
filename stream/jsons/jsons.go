package jsons

import (
	"fmt"

	"github.com/0xdbf/sigma"
	"github.com/0xdbf/sigma/core"
	"github.com/0xdbf/sigma/stream"
)

type Value map[string]interface{}

func (Value) Eq(core.Value) bool {
	return false
}

/*

Σ ...
*/
type σ string

/*

Σ ...
*/
func Σ() sigma.Σ {
	return σ("json")
}

func (σ) Stream(q core.Values) stream.Stream {
	fmt.Println(q)

	switch s := q.Get(core.Int(0)).(type) {
	case Value:
		switch p := q.Get(core.Int(1)).(type) {
		case core.Atom:
			switch v := s[string(p)].(type) {
			case map[string]interface{}:
				fmt.Println("map => ", v)
			case []interface{}:
				fmt.Println("seq => ", v)
			default:
				fmt.Println("scl => ", v)
			}
		}
	}

	return nil
}
