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

	"github.com/kshard/sigma/ast"
)

type Parser struct {
	s    *Scanner
	head *Token
	last *Token
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
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

// STRING ( TERMS ) :- BODY .
// STRING ( TERMS ) .
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

// STRING ( TERMS )
func (p *Parser) parseImply() (*ast.Imply, error) {
	imply := new(ast.Imply)

	token := p.scanSkipWhitespace()
	if token.Kind != STRING {
		return nil, fmt.Errorf("syntax error: Imply expect STRING, got %+v", token)
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

// STRING
// NUMBER
// " STRING "
func (p *Parser) parseTerm() (*ast.Term, error) {
	token := p.scanSkipWhitespace()
	switch {
	case token.Kind == STRING:
		return &ast.Term{Name: token.Literal}, nil
	// case token.Kind == NUMBER:
	// case token.Kind == SYMBOL && token.Literal == `"`:
	default:
		return nil, fmt.Errorf("syntax error: Term do not expect %+v", token)
	}
}
