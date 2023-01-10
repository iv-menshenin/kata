package topfreq

import "sort"

// We have an array of numbers in scattered and the number K,
// we need to choose K numbers with the greatest number of repetitions in this array.

func getTopFrequent(a []int64, k int) []int64 {
	var (
		result []int64
		counts = make(map[int64]int)
	)
	sort.Slice(a, func(i, j int) bool {
		return a[j] < a[i]
	})
	for _, num := range a {
		counts[num]++
		if len(result) == 0 || result[len(result)-1] != num {
			result = append(result, num)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return counts[result[j]] < counts[result[i]]
	})
	if len(result) > k {
		return result[:k]
	}
	return result
}
