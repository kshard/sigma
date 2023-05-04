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

package ast

//
// The file defines abstract syntax tree for expressing on σ-calculus
//

import (
	"fmt"

	"github.com/kshard/xsd"
)

// Kind of ast node
type Kind int

const (
	Invalid Kind = iota
	NodeTerm
	NodeFact
	NodeHorn
)

// Term defines complex term expressions of σ-expression
type Term struct {
	Name  string
	Value xsd.Value
}

func (*Term) Node() Kind { return NodeTerm }

func (x *Term) String() string {
	if x.Value != nil {
		return fmt.Sprintf("%v = '%v'", x.Name, x.Value)
	}

	return fmt.Sprintf("$%v", x.Name)
}

// Terms is the ordered set of terms
type Terms []*Term

// σ-expression
type Imply struct {
	Name  string
	Terms Terms
}

func (i *Imply) Term(term string, value ...xsd.Value) *Imply {
	t := &Term{Name: term}
	if len(value) != 0 {
		t.Value = value[0]
	}
	i.Terms = append(i.Terms, t)

	return i
}

func NewExpr(name string) *Imply {
	return &Imply{Name: name, Terms: make(Terms, 0)}
}

// Sequence of σ-expressions
type Implies []*Imply

// σ-expression
type Rule interface{ Rule() Kind }

type Rules []Rule

// σ-expression, Generator of ground facts
// TODO: deprecate type (see Head)
type Fact Imply

func (*Fact) Node() Kind { return NodeFact }
func (*Fact) Rule() Kind { return NodeFact }

// Helper configs Fact Node with Terms
func (f *Fact) Tuple(term ...string) *Fact {
	for _, t := range term {
		f.Terms = append(f.Terms, &Term{Name: t})
	}

	return f
}

// Helper instantiates Fact Node
func NewFact(name string) *Fact {
	return &Fact{Name: name, Terms: make(Terms, 0)}
}

// Head of Horn clause
type Head Imply

// Helper configs Fact Node with Terms
func (f *Head) Tuple(term ...string) *Head {
	f.Terms = make(Terms, len(term))
	for i, t := range term {
		f.Terms[i] = &Term{Name: t}
	}

	return f
}

func NewHead(name string) *Head {
	return &Head{Name: name, Terms: make(Terms, 0)}
}

// Horn clause is a syntax sugar for the projection (⟻) and join (⨝) operator
// It joins body and projects result as stream defined by the head.
//
//	H ⟻ A ⨝ B ⨝ ... ⨝ C
type Horn struct {
	Head *Head
	Body Implies
}

func (*Horn) Node() Kind { return NodeHorn }
func (*Horn) Rule() Kind { return NodeHorn }

func NewHorn(head *Head, imply ...*Imply) *Horn {
	return &Horn{Head: head, Body: imply}
}
