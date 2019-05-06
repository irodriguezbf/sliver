// Code generated by command: go run asm.go -out cast.s -stubs stub.go. DO NOT EDIT.

package cast

// Split returns the low 64, 32, 16 and 8 bits of x.
// Tests the As() methods of virtual general-purpose registers.
func Split(x uint64) (q uint64, l uint32, w uint16, b uint8)