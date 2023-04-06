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

package vm

import "fmt"

//
// The file defines heap address
//

// Addr is data type on defining heap addresses
type Addr uint32

const (
	fReadOnly = 1 << 31
	fAddrMask = 0x7fffffff
)

func (addr Addr) ReadOnly() Addr { return addr | fReadOnly }

func (addr Addr) IsWritable() bool { return addr&fReadOnly == 0 }
func (addr Addr) Value() uint32    { return (uint32(addr) & fAddrMask) }
func (addr Addr) String() string {
	f := "r"
	if addr.IsWritable() {
		f = "w"
	}

	return fmt.Sprintf("%s%d", f, addr.Value())
}
