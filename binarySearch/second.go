package main

import "sync/atomic"

// find2 will search for an integer value in an sorted array relying on an even distribution of values
//
// the complexity of the algorithm at best O(n), and at worst O(log n)
func find2(a []int, f int, iter ...*uint32) int {
	tick := func() {}
	if len(iter) > 0 {
		tick = func() {
			atomic.AddUint32(iter[0], 1)
		}
	}
	var (
		max     = len(a) - 1
		low, hi = 0, max
	)
	for {
		if hi < low {
			return -1
		}
		if hi == low {
			if a[hi] == f {
				return hi
			}
			return -1
		}
		tick()
		d := float64(a[hi]-a[low]) / float64(hi-low)
		mid := low + int(float64(f-a[low])/d)
		if mid < low || mid > hi {
			return -1
		}
		if a[mid] > f {
			if !(hi > mid-1) {
				return -1
			}
			hi = mid - 1
		} else if a[mid] < f {
			if !(low < mid+1) {
				return -1
			}
			low = mid + 1
		} else {
			return mid
		}
	}
}
