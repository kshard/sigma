package ast

//
type Kind int

const (
	Invalid Kind = iota
	NodeLit
	NodeVar
	NodeFact
	NodeHorn
)

//
type Term interface{ Term() Kind }

type Terms []Term

//
type Lit struct {
	ID    string
	Value any
}

func (*Lit) Term() Kind { return NodeLit }

//
type Var struct {
	Name string
}

func (*Var) Term() Kind { return NodeVar }

//
type Imply struct {
	Name  string
	Terms Terms
}

//
type Rule interface{ Rule() Kind }

type Rules []Rule

//
type Fact Imply

func (*Fact) Rule() Kind { return NodeFact }

//
type Horn struct {
	Head *Imply
	Body []*Imply
}

func (*Horn) Rule() Kind { return NodeHorn }
