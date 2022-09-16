package ast

import "fmt"

//
type Kind int

const (
	Invalid Kind = iota
	NodeTerm
	NodeFact
	NodeHorn
)

//
// type Term interface{ Term() Kind }

//
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

type Terms []*Term

//
// type Var struct {
// 	Name string
// }

// func (*Var) Term() Kind       { return NodeVar }
// func (x *Var) String() string { return fmt.Sprintf("$%v", x.Name) }

//
type Imply struct {
	Name  string
	Terms Terms
}

type Implies []*Imply

//
type Rule interface{ Rule() Kind }

type Rules []Rule

//
type Fact Imply

func (*Fact) Node() Kind { return NodeFact }
func (*Fact) Rule() Kind { return NodeFact }

//
type Horn struct {
	// TODO: Head Type where only Vars allowed
	Head *Imply
	Body Implies
}

func (*Horn) Node() Kind { return NodeHorn }
func (*Horn) Rule() Kind { return NodeHorn }
