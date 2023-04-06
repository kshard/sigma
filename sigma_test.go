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

package sigma_test

import (
	"reflect"
	"testing"

	"github.com/kshard/sigma"
	"github.com/kshard/sigma/ast"
	"github.com/kshard/sigma/internal/gen"
)

func queryMatchPerson() sigma.Reader {
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

	return sigma.New("h", rules)
}

func TestBasicQueryMatchPerson(t *testing.T) {
	sequence := queryMatchPerson().ToSeq()
	required := [][]any{{"urn:person:137"}}

	if !reflect.DeepEqual(sequence, required) {
		t.Errorf("got %v required %v", sequence, required)
	}
}

func BenchmarkBasicQueryMatchPerson(b *testing.B) {
	reader := queryMatchPerson()

	for i := 0; i < b.N; i++ {
		for {
			if err := reader.Read(nil); err != nil {
				break
			}
		}
	}
}

func queryMatchMovieByYear() sigma.Reader {
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

	return sigma.New("h", rules)
}

func TestBasicQueryMatchMovieByYear(t *testing.T) {

	sequence := queryMatchMovieByYear().ToSeq()
	required := [][]any{
		{"urn:movie:202", "Predator"},
		{"urn:movie:203", "Lethal Weapon"},
		{"urn:movie:204", "RoboCop"},
	}

	if !reflect.DeepEqual(sequence, required) {
		t.Errorf("got %v required %v", sequence, required)
	}
}

func BenchmarkBasicQueryMatchMovieByYear(b *testing.B) {
	reader := queryMatchMovieByYear()

	for i := 0; i < b.N; i++ {
		for {
			if err := reader.Read(nil); err != nil {
				break
			}
		}
	}
}

func queryDiscoverAllActorsFromMovie() sigma.Reader {
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

	return sigma.New("h", rules)
}

func TestBasicQueryDiscoverAllActorsFromMovie(t *testing.T) {
	sequence := queryDiscoverAllActorsFromMovie().ToSeq()
	required := [][]any{
		{"Mel Gibson"},
		{"Danny Glover"},
		{"Gary Busey"},
	}

	if !reflect.DeepEqual(sequence, required) {
		t.Errorf("got %v required %v", sequence, required)
	}
}

func BenchmarkBasicQueryDiscoverAllActorsFromMovie(b *testing.B) {
	reader := queryDiscoverAllActorsFromMovie()

	for i := 0; i < b.N; i++ {
		for {
			if err := reader.Read(nil); err != nil {
				break
			}
		}
	}
}

func BenchmarkCompileBasicQueryDiscoverAllActorsFromMovie(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queryDiscoverAllActorsFromMovie()
	}
}
