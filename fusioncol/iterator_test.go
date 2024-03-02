package fusioncol

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIterator(t *testing.T) {
	t.Parallel()
	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		var c Collection[int]
		for x := range c.Iterator(context.Background(), false, 16) {
			t.Errorf("unexpected: %d", x)
		}
	})
	t.Run("Forward", func(t *testing.T) {
		t.Parallel()
		var c Collection[int]
		const max = 100000
		for n := 0; n < max; n++ {
			c.Append(n)
		}
		var i int
		for x := range c.Iterator(context.Background(), false, 16) {
			require.Equal(t, i, *x)
			i++
		}
	})
	t.Run("Backward", func(t *testing.T) {
		t.Parallel()
		var c Collection[int]
		const max = 100000
		for n := 0; n < max; n++ {
			c.Append(n)
		}
		var i = max
		for x := range c.Iterator(context.Background(), true, 16) {
			i--
			require.Equal(t, i, *x)
		}
	})
	t.Run("Cancel", func(t *testing.T) {
		t.Parallel()
		var c Collection[int]
		for n := 0; n < 10; n++ {
			c.Append(0)
		}
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		ch := c.Iterator(ctx, false, 0)
		<-ch
		ch = c.Iterator(ctx, true, 0)
		<-ch
	})
}

func BenchmarkIterator(b *testing.B) {
	type Elem struct {
		s          string
		a, b, c, d int64
		n          int
	}
	var c Collection[Elem]
	const max = 1000000
	for n := 0; n < max; n++ {
		c.Append(Elem{n: n})
	}
	b.Run("Forward", func(b *testing.B) {
		b.ReportAllocs()
		for x := range c.Iterator(context.Background(), false, 16) {
			x = x
		}
	})
	b.Run("Backward", func(b *testing.B) {
		b.ReportAllocs()
		for x := range c.Iterator(context.Background(), true, 16) {
			x = x
		}
	})
}
