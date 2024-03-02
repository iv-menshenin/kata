package fusioncol

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortable(t *testing.T) {
	t.Parallel()
	t.Run("Sort10", func(t *testing.T) {
		t.Parallel()
		var c Collection[int]
		s := Sortable(&c, func(i *int, j *int) bool {
			return *i < *j
		})
		c.Append(0)
		c.Append(2)
		c.Append(6)
		c.Append(3)
		c.Append(5)
		c.Append(8)
		c.Append(9)
		c.Append(4)
		c.Append(1)
		c.Append(7)
		sort.Sort(s)
		for n := 0; n < 10; n++ {
			require.Equal(t, n, *c.Get(n))
		}
	})
	t.Run("Sort10", func(t *testing.T) {
		t.Parallel()
		var c Collection[int]
		s := Sortable(&c, func(i *int, j *int) bool {
			return *i < *j
		})
		const max = 1000000
		for n := max; n > 0; n-- {
			c.Append(n - 1)
		}
		sort.Sort(s)
		for n := 0; n < max; n++ {
			require.Equal(t, n, *c.Get(n))
		}
	})
}

func BenchmarkSortable(b *testing.B) {
	var c Collection[string]
	const max = 1000
	for n := 0; n < max; n++ {
		c.Append(strconv.Itoa(n))
	}
	s := Sortable(&c, func(i *string, j *string) bool {
		return *i < *j
	})
	b.Run("Sortable", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			if n%2 == 0 {
				sort.Sort(s)
			} else {
				sort.Sort(sort.Reverse(s))
			}
		}
	})
	b.Run("IsSorted", func(b *testing.B) {
		sort.Sort(s)
		b.ResetTimer()
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			sort.IsSorted(s)
		}
	})
	b.Run("NotSorted", func(b *testing.B) {
		if b.N > 1 {
			s.Swap(0, 1)
		}
		b.ResetTimer()
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			sort.IsSorted(s)
		}
	})
}
