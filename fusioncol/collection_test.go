package fusioncol

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFusionCollection(t *testing.T) {
	t.Parallel()
	t.Run("append_get", func(t *testing.T) {
		t.Parallel()

		var c Collection[int64]
		c.Append(101)
		c.Append(102)
		c.Append(103)
		require.Equal(t, int64(103), *c.Get(2))
		require.Equal(t, int64(102), *c.Get(1))
		require.Equal(t, int64(101), *c.Get(0))
	})
	t.Run("len", func(t *testing.T) {
		t.Parallel()

		var c Collection[int64]
		require.Equal(t, 0, c.Len())
		c.Append(101)
		require.Equal(t, 1, c.Len())
		c.Append(102)
		require.Equal(t, 2, c.Len())
		c.Append(103)
		require.Equal(t, 3, c.Len())
	})
	t.Run("force", func(t *testing.T) {
		t.Parallel()
		type Elem struct {
			i int
			s string
		}
		var c Collection[Elem]

		const elemCount = 100000

		for n := 0; n < elemCount; n++ {
			require.Equal(t, n, c.Len())
			c.Append(Elem{i: n, s: strconv.Itoa(n)})
		}

		for n := 0; n < elemCount; n++ {
			e := c.Get(n)
			require.Equal(t, n, e.i)
			require.Equal(t, strconv.Itoa(n), e.s)
		}
	})
}

func TestFusionCollectionPushPop(t *testing.T) {
	t.Parallel()
	t.Run("push_pop", func(t *testing.T) {
		t.Parallel()
		var c Collection[string]

		c.Push("foo")
		require.Equal(t, 1, c.Len())
		c.Push("bar")
		require.Equal(t, 2, c.Len())
		require.Equal(t, "bar", c.Pop())
		require.Equal(t, 1, c.Len())
		require.Equal(t, "foo", c.Pop())
		require.Equal(t, 0, c.Len())

		c.Push("1")
		c.Push("3")
		require.Equal(t, "3", c.Pop())
		c.Push("2")
		c.Push("3")
		c.Push("4")
		require.Equal(t, 4, c.Len())
		require.Equal(t, "4", c.Pop())
		require.Equal(t, "3", c.Pop())
		require.Equal(t, "2", c.Pop())
		require.Equal(t, "1", c.Pop())
		require.Equal(t, 0, c.Len())
	})
	t.Run("force", func(t *testing.T) {
		t.Parallel()

		var c Collection[int]
		const elemCount = 1000000

		for n := 0; n < elemCount; n++ {
			c.Push(n)
		}
		for n := elemCount - 1; n >= 0; n-- {
			require.Equal(t, n, c.Pop())
		}
	})
}

func BenchmarkFusionCollectionAppendGet(b *testing.B) {
	b.ReportAllocs()
	type Elem struct {
		s          string
		a, b, c, d int64
		n          int
	}
	var c Collection[Elem]
	for n := 0; n < b.N; n++ {
		c.Append(Elem{n: n})
	}
	for n := 0; n < b.N; n++ {
		_ = c.Get(n)
	}
}

func BenchmarkFusionCollectionPushPop(b *testing.B) {
	b.ReportAllocs()
	type Elem struct {
		s          string
		a, b, c, d int64
		n          int
	}
	var c Collection[Elem]
	for n := 0; n < b.N; n++ {
		c.Push(Elem{n: n})
	}
	for n := 0; n < b.N; n++ {
		_ = c.Pop()
	}
}
