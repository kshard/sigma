package lang_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/kshard/sigma"
	"github.com/kshard/sigma/asm"
	"github.com/kshard/sigma/internal/gen"
	"github.com/kshard/sigma/lang"
)

func TestXxx(t *testing.T) {
	q := `
		f(s, p, o).

		a(movie, cast) :-
			f(m, "title", movie),
			f(m, "cast", cast).

		h(name, eman) :-
			a("Lethal Weapon", p),
			f(p, "name", name),

			a("Mad Max", s),
			f(s, "name", eman).
	`

	buf := bytes.NewBuffer([]byte(q))

	parser := lang.NewParser(buf)
	rules, err := parser.Parse()
	fmt.Printf("%s, %+v\n", err, rules)

	machine, err := sigma.New("h", rules)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", machine.Code)

	ctx := asm.NewContext().Add("f", gen.FactsIMDB)
	reader := sigma.Stream(ctx, machine)
	fmt.Println(reader.ToSeq())

}
