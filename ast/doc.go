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

// Package ast defines abstract syntax tree for expressing on σ-calculus
//
// Examples on the usage of ast
// Generator is a σ-expressions producing stream of binary relations
// imdb(s, p, o) ⇒ ⟨subject, predicate, object⟩.
//
//	&ast.Fact{Stream: }
package ast
