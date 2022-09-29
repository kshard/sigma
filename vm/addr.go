/*

  Copyright 2016 Dmitry Kolesnikov, All Rights Reserved

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package vm

import "fmt"

//
// The file defines heap address
//

/*

Addr is data type on defining heap addresses
*/
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
