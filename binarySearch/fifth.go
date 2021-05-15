package main

import "sync/atomic"

// find5 will search for an integer value in an sorted array.
// implemented using a closure inside a function
//
// the complexity of the algorithm at best O(n), and at worst O(log n)
func find5(a []int, i int, iter ...*uint32) int {
	var findInternal func([]int, int) int
	findInternal = func(a []int, offset int) int {
		if len(iter) > 0 {
			atomic.AddUint32(iter[0], 1)
		}
		if len(a) < 2 {
			if len(a) == 1 && a[0] == i {
				return offset
			}
			return -1
		}
		var (
			hi  = len(a) - 1
			del = float64(a[hi]-a[0]) / float64(hi)
			mid = int(float64(i-a[0]) / del)
		)
		if mid < 0 || mid > hi {
			return -1
		}
		if a[mid] == i {
			return mid + offset
		}
		if a[mid] < i {
			return findInternal(a[mid+1:], offset+mid+1)
		} else {
			return findInternal(a[:mid], offset)
		}
	}
	return findInternal(a, 0)
}
