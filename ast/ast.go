/*

  Copyright 2016 Dmitry Kolesnikov, All Rights Reserved

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

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
