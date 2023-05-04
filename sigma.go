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

package sigma

import (
	"github.com/kshard/sigma/asm"
	"github.com/kshard/sigma/ast"
	"github.com/kshard/sigma/internal/compiler"
	"github.com/kshard/sigma/vm"
	"github.com/kshard/xsd"
)

//
// The file defines public api for Sigma VM
//

// Reader is a stream produced by evaluation of rules
type Reader interface {
	ToSeq() [][]xsd.Value
	Read([]xsd.Value) error
}

type VM struct {
	Machine *vm.VM
	Shape   []vm.Addr
	Code    asm.Stream
}

// Create instance of VM
func New(goal string, rules ast.Rules) (*VM, error) {
	sc := compiler.New()
	if err := sc.Compile(rules); err != nil {
		return nil, err
	}

	machine, shape, code := sc.Assemble(goal)

	return &VM{
		Machine: machine,
		Shape:   shape,
		Code:    code,
	}, nil
}

func Stream(ctx *asm.Context, vm *VM) Reader {
	stream := vm.Code.Link(ctx)
	return vm.Machine.Stream(vm.Shape, stream)
}

func NewReader(ctx *asm.Context, goal string, rules ast.Rules) (Reader, error) {
	machine, err := New(goal, rules)
	if err != nil {
		return nil, err
	}

	return Stream(ctx, machine), nil
}
