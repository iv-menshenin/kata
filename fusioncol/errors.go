package fusioncol

import "fmt"

type ErrOutOfBounds struct {
	l, i int
}

func (e ErrOutOfBounds) Error() string {
	return fmt.Sprintf("index out of bounds: requested %d element with %d length", e.i, e.l)
}
