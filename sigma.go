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

package sigma

import (
	"github.com/kshard/sigma/ast"
	"github.com/kshard/sigma/internal/compile"
)

//
// The file defines public api for Sigma VM
//

// Reader is a stream produced by evaluation of rules
type Reader interface {
	ToSeq() [][]any
	Read([]any) error
}

// New creates a new Reader
func New(goal string, rules ast.Rules) Reader {
	c := compile.New()
	c.Compile(rules)
	// vmm, addr, reader := c.ReaderX(goal)
	// return vmm.Stream(addr, reader.Compile())
	return nil
}
