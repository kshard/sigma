/*

  Sigma Virtual Machine
  Copyright (C) 2016 - 2023  Dmitry Kolesnikov

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
	"bytes"
	"encoding/gob"

	"github.com/kshard/sigma/asm"
	"github.com/kshard/sigma/vm"
)

func init() {
	var (
		tstring string
		astring any = &tstring
		tint    int
		aint    any = &tint
	)

	gob.Register(&asm.Generator{})
	gob.Register(&asm.Horn{})
	gob.Register(&vm.VM{})
	gob.Register(&astring)
	gob.Register(&aint)
}

func Encode(vm *VM) ([]byte, error) {
	var buf bytes.Buffer
	codec := gob.NewEncoder(&buf)

	if err := codec.Encode(vm); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Decode(buf []byte) (*VM, error) {
	var vm VM
	codec := gob.NewDecoder(bytes.NewBuffer(buf))

	if err := codec.Decode(&vm); err != nil {
		return nil, err
	}

	return &vm, nil
}
