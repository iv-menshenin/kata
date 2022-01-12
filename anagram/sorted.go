package anagram

import "sort"

type intSorted []rune

func (s intSorted) Len() int {
	return len(s)
}

func (s intSorted) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s intSorted) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func isAnagramSorted(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return false
	}
	var (
		aI = []rune(a)
		bI = []rune(b)
	)
	sort.Sort(intSorted(aI))
	sort.Sort(intSorted(bI))
	for i := range aI {
		if aI[i] != bI[i] {
			return false
		}
	}
	return true
}
