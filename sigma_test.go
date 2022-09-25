package sigma_test

import (
	"reflect"
	"testing"

	"github.com/0xdbf/sigma"
	"github.com/0xdbf/sigma/ast"
	"github.com/0xdbf/sigma/internal/gen"
)

func TestBasicQueryMatchPerson(t *testing.T) {
	rules := ast.Rules{
		&ast.Fact{
			Stream:    &ast.Imply{Name: "f", Terms: ast.Terms{{Name: "s"}, {Name: "p"}, {Name: "o"}}},
			Generator: gen.FactsIMDB,
		},

		&ast.Horn{
			Head: &ast.Head{Name: "h", Terms: ast.Terms{{Name: "s"}}},
			Body: []*ast.Imply{
				{Name: "f", Terms: ast.Terms{
					{Name: "s"},
					{Name: "t1", Value: "name"},
					{Name: "t2", Value: "Ridley Scott"},
				}},
			},
		},
	}

	sequence := sigma.New("h", rules).ToSeq()
	required := [][]any{{"urn:person:137"}}

	if !reflect.DeepEqual(sequence, required) {
		t.Errorf("got %v required %v", sequence, required)
	}
}

func TestBasicQueryMatchMovieByYear(t *testing.T) {
	rules := ast.Rules{
		&ast.Fact{
			Stream:    &ast.Imply{Name: "f", Terms: ast.Terms{{Name: "s"}, {Name: "p"}, {Name: "o"}}},
			Generator: gen.FactsIMDB,
		},

		&ast.Horn{
			Head: &ast.Head{Name: "h", Terms: ast.Terms{{Name: "s"}, {Name: "title"}}},
			Body: []*ast.Imply{
				{Name: "f", Terms: ast.Terms{
					{Name: "s"},
					{Name: "t1", Value: "year"},
					{Name: "t2", Value: 1987},
				}},
				{Name: "f", Terms: ast.Terms{
					{Name: "s"},
					{Name: "t3", Value: "title"},
					{Name: "title"},
				}},
			},
		},
	}

	sequence := sigma.New("h", rules).ToSeq()
	required := [][]any{
		{"urn:movie:202", "Predator"},
		{"urn:movie:203", "Lethal Weapon"},
		{"urn:movie:204", "RoboCop"},
	}

	if !reflect.DeepEqual(sequence, required) {
		t.Errorf("got %v required %v", sequence, required)
	}
}

func TestBasicQueryDiscoverAllActorsFromMovie(t *testing.T) {
	rules := ast.Rules{
		&ast.Fact{
			Stream:    &ast.Imply{Name: "f", Terms: ast.Terms{{Name: "s"}, {Name: "p"}, {Name: "o"}}},
			Generator: gen.FactsIMDB,
		},

		&ast.Horn{
			Head: &ast.Head{Name: "h", Terms: ast.Terms{{Name: "name"}}},
			Body: []*ast.Imply{
				{Name: "f", Terms: ast.Terms{
					{Name: "m"},
					{Name: "t1", Value: "title"},
					{Name: "t2", Value: "Lethal Weapon"},
				}},
				{Name: "f", Terms: ast.Terms{
					{Name: "m"},
					{Name: "t3", Value: "cast"},
					{Name: "p"},
				}},
				{Name: "f", Terms: ast.Terms{
					{Name: "p"},
					{Name: "t4", Value: "name"},
					{Name: "name"},
				}},
			},
		},
	}

	sequence := sigma.New("h", rules).ToSeq()
	required := [][]any{
		{"Mel Gibson"},
		{"Danny Glover"},
		{"Gary Busey"},
	}

	if !reflect.DeepEqual(sequence, required) {
		t.Errorf("got %v required %v", sequence, required)
	}
}
