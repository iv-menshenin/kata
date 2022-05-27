package binaryVector

// You just need to count the number of maximum number of consecutive signals 1

func countSerials(v []int8) int {
	var max int
	var curr int
	for _, bit := range v {
		if bit == 1 {
			if curr++; curr > max {
				max = curr
			}
		} else {
			curr = 0
		}
	}
	return max
}
