package topfreq

import "sort"

// We have an array of numbers in scattered and the number K,
// we need to choose K numbers with the greatest number of repetitions in this array.

func getTopFrequent(a []int64, k int) []int64 {
	var (
		result []int64
		counts []int
	)
	sort.Slice(a, func(i, j int) bool {
		return a[j] < a[i]
	})
	for _, num := range a {
		if len(result) == 0 || result[len(result)-1] != num {
			result = append(result, num)
			counts = append(counts, 1)
			continue
		}
		counts[len(counts)-1]++
	}
	sort.Slice(result, func(i, j int) bool {
		return counts[j] < counts[i]
	})
	if len(result) > k {
		return result[:k]
	}
	return result
}
