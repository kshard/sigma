package lang_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/kshard/sigma/lang"
)

func TestXxx(t *testing.T) {
	buf := bytes.NewBuffer([]byte("h(s,p,o) :- f(s,p,o)."))

	parser := lang.NewParser(buf)
	rules, err := parser.Parse()

	fmt.Printf("%s, %+v\n", err, rules)
}
