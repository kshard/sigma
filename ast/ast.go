/*

  Sigma Virtual Machine
  Copyright (C) 2016  Dmitry Kolesnikov

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

	"github.com/kshard/sigma/vm"
)

//
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
	Value any
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

//
type Implies []*Imply

//
type Rule interface{ Rule() Kind }

type Rules []Rule

//
type Fact struct {
	Stream    *Imply
	Generator func(Terms) vm.Generator
}

func (*Fact) Node() Kind { return NodeFact }
func (*Fact) Rule() Kind { return NodeFact }

//
type Head Imply

//
type Horn struct {
	Head *Head
	Body Implies
}

func (*Horn) Node() Kind { return NodeHorn }
func (*Horn) Rule() Kind { return NodeHorn }
