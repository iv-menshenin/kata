package binaryVector

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
