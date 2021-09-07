package sigma

import (
	"github.com/0xdbf/sigma/core"
	"github.com/0xdbf/sigma/stream"
)

/*

Σ .... (sigma function implemented by data source / generator)
*/
type Σ interface {
	Stream(q core.Values) stream.Stream
}
