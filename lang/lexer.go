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
	"bufio"
	"bytes"
	"io"
)

type Token struct {
	Kind    Kind
	Literal string
}

type Kind int

const (
	UNKNOWN Kind = iota
	EOF
	WS

	IF

	SYMBOL
	ATOM
	STRING
	NUMBER
	DECIMAL
)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isSymbol(ch rune) bool {
	return ch == ':' || ch == '-' || ch == '(' || ch == ',' || ch == ')' || ch == '.'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isNumber(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

const eof = rune(0)

// xxx
type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

func (s *Scanner) Scan() Token {
	ch := s.read()

	switch {
	case isWhitespace(ch):
		s.unread()
		return s.scanWhitespace()

	case isLetter(ch):
		s.unread()
		return s.scanAtom()

	case ch == '"':
		s.unread()
		return s.scanString()

	case isNumber(ch):
		s.unread()
		return s.scanNumber()

	case ch == ':':
		if ch = s.read(); ch == '-' {
			return Token{Kind: IF}
		} else {
			s.unread()
			return Token{Kind: SYMBOL, Literal: ":"}
		}

	case isSymbol(ch):
		return Token{Kind: SYMBOL, Literal: string(ch)}

	case ch == eof:
		return Token{Kind: EOF}

	default:
		return Token{Kind: UNKNOWN, Literal: string(ch)}
	}
}

func (s *Scanner) scanWhitespace() Token {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return Token{Kind: WS, Literal: buf.String()}
}

func (s *Scanner) scanAtom() Token {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isNumber(ch) && ch != '_' {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return Token{Kind: ATOM, Literal: buf.String()}
}

func (s *Scanner) scanString() Token {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == '"' {
			buf.WriteRune(ch)
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return Token{Kind: STRING, Literal: buf.String()}
}

func (s *Scanner) scanNumber() Token {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isNumber(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	if ch := s.read(); ch != '.' {
		s.unread()
		return Token{Kind: NUMBER, Literal: buf.String()}
	} else {
		buf.WriteRune(ch)
	}

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isNumber(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return Token{Kind: DECIMAL, Literal: buf.String()}
}
