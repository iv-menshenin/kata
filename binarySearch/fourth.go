package main

import (
	"math"
	"sync/atomic"
)

// find1 will search for an integer value in an sorted array.
// implementation of the algorithm is simplified by using tail recursion
func find4(a []int, i int) int {
	return find4internal(a, i, 0)
}

func find4internal(a []int, i, offset int, iter ...*uint32) int {
	if len(a) == 0 {
		return -1
	}
	var (
		curr = 0
		prev = 0
		maxi = len(a) - 1
		step = int(math.Sqrt(float64(len(a))))
	)
	for {
		if len(iter) > 0 {
			atomic.AddUint32(iter[0], 1)
		}
		if a[curr] == i {
			return curr + offset
		}
		if a[curr] < i && curr < maxi {
			prev = curr
			if curr += step; curr > maxi {
				curr = maxi
			}
			step *= 2
		} else {
			break
		}
	}
	if curr == prev {
		return -1
	} else {
		a = a[prev+1 : curr]
		return find4internal(a, i, offset+prev+1, iter...)
	}
}
