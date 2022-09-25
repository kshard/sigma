package vm

import "fmt"

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
