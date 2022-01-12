package anagram

import "testing"

func Test_isAnagramSorted(t *testing.T) {
	testSomeAnagramChecker(t, isAnagramSorted)
}

func Benchmark_isAnagramSorted(t *testing.B) {
	benchSomeAnagramChecker(t, isAnagramSorted)
}
