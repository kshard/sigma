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
	"github.com/kshard/sigma/asm"
	"github.com/kshard/sigma/ast"
	"github.com/kshard/sigma/internal/gen"
)

func queryMatchPerson() sigma.Reader {
	rules := ast.Rules{
		ast.NewFact("f").Tuple("s", "p", "o"),
		ast.NewHorn(
			ast.NewHead("h").Tuple("s"),
			ast.NewExpr("f").
				Term("s").
				Term("t1", "name").
				Term("t2", "Ridley Scott"),
		),
	}

	reader, err := sigma.NewReader(asm.NewContext().Add("f", gen.FactsIMDB), "h", rules)
	if err != nil {
		panic(err)
	}

	return reader
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
		ast.NewFact("f").Tuple("s", "p", "o"),

		ast.NewHorn(
			ast.NewHead("h").Tuple("s", "title"),
			ast.NewExpr("f").
				Term("s").
				Term("t1", "year").
				Term("t2", 1987),
			ast.NewExpr("f").
				Term("s").
				Term("t3", "title").
				Term("title"),
		),
	}

	reader, err := sigma.NewReader(asm.NewContext().Add("f", gen.FactsIMDB), "h", rules)
	if err != nil {
		panic(err)
	}

	return reader
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
		ast.NewFact("f").Tuple("s", "p", "o"),

		ast.NewHorn(
			ast.NewHead("h").Tuple("name"),
			ast.NewExpr("f").
				Term("m").
				Term("t1", "title").
				Term("t2", "Lethal Weapon"),
			ast.NewExpr("f").
				Term("m").
				Term("t3", "cast").
				Term("p"),
			ast.NewExpr("f").
				Term("p").
				Term("t4", "name").
				Term("name"),
		),
	}

	reader, err := sigma.NewReader(asm.NewContext().Add("f", gen.FactsIMDB), "h", rules)
	if err != nil {
		panic(err)
	}

	return reader
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
