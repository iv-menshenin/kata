package optimization

import (
	"bytes"
)

//go:noinline
func strLenBytesBuffer() int {
	// the slice created inside the structure will be placed inside the heap
	var b bytes.Buffer
	b.WriteString("1")
	return b.Len()
}

//go:noinline
func strLenNewBufferString() int {
	// all structures will remain on the stack
	var b = bytes.NewBufferString("") // or NewBuffer(make([]byte, 0, 64))
	b.WriteString("make it easy")
	return b.Len()
}
