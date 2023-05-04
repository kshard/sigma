/*

  Sigma Virtual Machine
  Copyright (C) 2016 - 2023 Dmitry Kolesnikov

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as published
  by the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/

package lang

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/fogfish/curie"
	"github.com/kshard/sigma/ast"
	"github.com/kshard/xsd"
)

type Parser struct {
	s     *Scanner
	index int
	head  *Token
	last  *Token
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r), index: 0}
}

func (p *Parser) scan() (token Token) {
	if p.last != nil {
		token, p.last = *p.last, nil
		return
	}

	token = p.s.Scan()
	p.head = &token

	return
}

func (p *Parser) unscan() { p.last = p.head }

func (p *Parser) scanSkipWhitespace() (token Token) {
	token = p.scan()
	if token.Kind == WS {
		token = p.scan()
	}

	return
}

func (p *Parser) Parse() (ast.Rules, error) {
	rules := ast.Rules{}

	for {
		token := p.scanSkipWhitespace()
		if token.Kind == EOF {
			return rules, nil
		}

		p.unscan()

		rule, err := p.parseRule()
		if err != nil {
			return nil, err
		}

		rules = append(rules, rule)
	}
}

// ATOM ( TERMS ) :- BODY .
// ATOM ( TERMS ) .
func (p *Parser) parseRule() (ast.Rule, error) {
	imply, err := p.parseImply()
	if err != nil {
		return nil, err
	}

	token := p.scanSkipWhitespace()
	switch {
	case token.Kind == SYMBOL && token.Literal == `.`:
		fact := ast.Fact(*imply)
		return &fact, nil

	case token.Kind == IF:
		implies, err := p.parseImplies()
		if err != nil {
			return nil, err
		}

		token := p.scanSkipWhitespace()
		if token.Kind != SYMBOL && token.Literal != `.` {
			return nil, fmt.Errorf("syntax error: Rule expect `.`, got %+v", token)
		}

		head := ast.Head(*imply)
		horn := &ast.Horn{
			Head: &head,
			Body: implies,
		}
		return horn, nil

	default:
		return nil, fmt.Errorf("syntax error: Rule expect `.` or `:-`, got %+v", token)
	}
}

// ATOM ( TERMS ) , BODY
func (p *Parser) parseImplies() (ast.Implies, error) {
	implies := make(ast.Implies, 0)

	for {
		imply, err := p.parseImply()
		if err != nil {
			return nil, err
		}

		implies = append(implies, imply)

		token := p.scanSkipWhitespace()
		if token.Kind != SYMBOL || token.Literal != `,` {
			p.unscan()
			return implies, nil
		}
	}
}

// ATOM ( TERMS )
func (p *Parser) parseImply() (*ast.Imply, error) {
	imply := new(ast.Imply)

	token := p.scanSkipWhitespace()
	if token.Kind != ATOM {
		return nil, fmt.Errorf("syntax error: Imply expect ATOM, got %+v", token)
	}

	imply.Name = token.Literal

	token = p.scanSkipWhitespace()
	if token.Kind != SYMBOL && token.Literal != `(` {
		return nil, fmt.Errorf("syntax error: Imply expect `(`, got %+v", token)
	}

	terms, err := p.parseTerms()
	if err != nil {
		return nil, err
	}

	imply.Terms = terms

	token = p.scanSkipWhitespace()
	if token.Kind != SYMBOL && token.Literal != `)` {
		return nil, fmt.Errorf("syntax error: Imply expect `)`, got %+v", token)
	}

	return imply, nil
}

// TERM , TERMS
func (p *Parser) parseTerms() (ast.Terms, error) {
	terms := make(ast.Terms, 0)

	for {
		term, err := p.parseTerm()
		if err != nil {
			return nil, err
		}

		terms = append(terms, term)

		token := p.scanSkipWhitespace()
		if token.Kind != SYMBOL || token.Literal != `,` {
			p.unscan()
			return terms, nil
		}
	}
}

// ATOM
// < IRI >
// NUMBER
// DECIMAL
// " STRING "
func (p *Parser) parseTerm() (*ast.Term, error) {
	token := p.scanSkipWhitespace()
	switch {
	case token.Kind == ATOM:
		return &ast.Term{Name: token.Literal}, nil

	case token.Kind == XSD_ANYURI:
		name := "xu" + strconv.Itoa(p.index)
		p.index++
		iri := curie.IRI(strings.Trim(token.Literal, `<>`))
		return &ast.Term{Name: name, Value: xsd.ToAnyURI(iri)}, nil

	case token.Kind == XSD_STRING:
		name := "xs" + strconv.Itoa(p.index)
		p.index++
		return &ast.Term{Name: name, Value: xsd.From(strings.Trim(token.Literal, `"`))}, nil

	case token.Kind == XSD_INTEGER:
		name := "xn" + strconv.Itoa(p.index)
		p.index++

		val, err := strconv.Atoi(token.Literal)
		if err != nil {
			return nil, fmt.Errorf("syntax error: invalid int %+v, %s", token, err)
		}

		return &ast.Term{Name: name, Value: xsd.From(val)}, nil

	case token.Kind == XSD_DECIMAL:
		name := "xf" + strconv.Itoa(p.index)
		p.index++

		val, err := strconv.ParseFloat(token.Literal, 64)
		if err != nil {
			return nil, fmt.Errorf("syntax error: invalid float %+v, %s", token, err)
		}

		return &ast.Term{Name: name, Value: xsd.From(val)}, nil

	default:
		return nil, fmt.Errorf("syntax error: Term do not expect %+v", token)
	}
}
