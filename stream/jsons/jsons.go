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

func (sig σ) Stream(q core.Values) stream.Stream {
	switch pat := q.(type) {
	case core.Seq:
		return sig.stream(pat)
	default:
		// Note: internal design of VM uses core.Seq
		panic(fmt.Errorf("Invalid query type, core.Seq is expected"))
	}
}

func (sig σ) stream(q core.Seq) stream.Stream {
	switch subj := q[0].(type) {
	case Value:
		return sig.querySubject(subj, q)
	default:
		// Note: subject is always required at this design
		panic(fmt.Errorf("Subject is not defined"))
	}
}

func (sig σ) querySubject(subj Value, q core.Seq) stream.Stream {
	switch pred := q[1].(type) {
	case core.Atom:
		return sig.querySubjectPredicate(subj, pred, q)
	default:
		// Note: only expact match of predicate is supported
		panic(fmt.Errorf("Unsupported predicate query type %T", pred))
	}
}

func (sig σ) querySubjectPredicate(subj Value, pred core.Atom, q core.Seq) stream.Stream {
	switch obj := q[2].(type) {
	case core.Atom:
		if obj == core.Any {
			return sig.returnObject(subj, pred, q)
		}
		return nil
	default:
		// Note: only expact match of predicate is supported
		panic(fmt.Errorf("Unsupported object query type %T", pred))
	}
}

func (σ) returnObject(subj Value, pred core.Atom, q core.Seq) stream.Stream {
	switch obj := subj[string(pred)].(type) {
	case map[string]interface{}:
		q[2] = Value(obj)
		return stream.Unit(q)
	case []interface{}:
		fmt.Println("seq => ", obj)
		return nil
	case float64:
		q[2] = core.Float64(obj)
		return stream.Unit(q)
	default:
		fmt.Println("scl => ", obj)
		return nil
	}
}
