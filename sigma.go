package sigma

import (
	"github.com/0xdbf/sigma/ast"
	"github.com/0xdbf/sigma/internal/compile"
)

type Reader interface {
	ToSeq() [][]any
	Read([]any) error
}

func New(goal string, rules ast.Rules) Reader {
	c := compile.New()
	c.Compile(rules)
	return c.Reader(goal)
}
