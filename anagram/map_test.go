package anagram

import "testing"

func Test_isAnagramMapped(t *testing.T) {
	testSomeAnagramChecker(t, isAnagramMapped)
}

func Benchmark_isAnagramMapped(t *testing.B) {
	benchSomeAnagramChecker(t, isAnagramMapped)
}
