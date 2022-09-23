package ast

//
// The file defines abstract syntax tree for expressing on σ-calculus
//

import "fmt"

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

// Terms ordered set
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
type Fact Imply

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
