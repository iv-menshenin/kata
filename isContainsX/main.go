package isContainsX

import "sort"

// i != j
// 0 <= i, j < arr.length
// arr[i] == 2 * arr[j]
// 2 <= arr.length <= 500
// 0 <= arr[i] <= 10^3

func byArraySlow(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] == 2*arr[j] {
				return true
			}
		}
	}
	return false
}

func byMap(arr []int) bool {
	sort.Ints(arr)
	var d = make(map[int]struct{}, len(arr))
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			var divided = arr[i] / 2
			if _, ok := d[divided]; ok {
				return ok
			}
		}
		d[arr[i]] = struct{}{}
	}
	return false
}

func byMap2(arr []int) bool {
	var ia = make(map[int]struct{}, len(arr))
	for _, i := range arr {
		if _, ok := ia[i*2]; ok {
			return true
		}
		if i%2 == 0 {
			if _, ok := ia[i/2]; ok {
				return true
			}
		}
		ia[i] = struct{}{}
	}
	return false
}

func byIntArr(arr []int) bool {
	var ia = make([]int, 1001)
	for _, i := range arr {
		if ia[i] == 2 {
			return true
		}
		ia[i] = 1
		if i <= 500 {
			ia[i*2] = 2
		}
		if i%2 == 0 {
			ia[i/2] = 2
		}
	}
	return false
}
