package main

import "sync/atomic"

// find1 will search for an integer value in an sorted array.
// implementation of the algorithm is simplified by using tail recursion
func find3(a []int, i int) int {
	return find3internal(a, i, 0)
}

func find3internal(a []int, i, offset int, iter ...*uint32) int {
	if len(iter) > 0 {
		atomic.AddUint32(iter[0], 1)
	}
	if len(a) == 0 {
		return -1
	}
	var mid = len(a) / 2
	if a[mid] == i {
		return mid + offset
	}
	if a[mid] < i {
		return find3internal(a[mid+1:], i, offset+mid+1, iter...)
	} else {
		return find3internal(a[:mid], i, offset, iter...)
	}
}
