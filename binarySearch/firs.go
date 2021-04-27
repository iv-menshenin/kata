package main

import "sync/atomic"

func find1(a []int, f int, iter ...*uint32) int {
	tick := func() {}
	if len(iter) > 0 {
		tick = func() {
			atomic.AddUint32(iter[0], 1)
		}
	}
	var (
		max = len(a) - 1
		low = 0
		hi  = max
	)
	for {
		tick()
		mid := low + (hi-low)/2
		if mid < 0 || mid > max {
			return -1
		}
		if a[mid] < f {
			if low == mid+1 {
				return -1
			}
			low = mid + 1
		} else if a[mid] > f {
			if hi == mid-1 {
				return -1
			}
			hi = mid - 1
		} else {
			return mid
		}
	}
}
