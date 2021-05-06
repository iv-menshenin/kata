package main

import (
	"testing"
)

func Test_find3(t *testing.T) {
	goTestFind(t, func(a []int, b int) int {
		return find3(a, b)
	})
}

func Test_find3_brutForce(t *testing.T) {
	goTestBrutForce(t, func(a []int, b int, c ...*uint32) int {
		return find3internal(a, b, 0, c...)
	})
}

func Test_find3_similarValues(t *testing.T) {
	goTestSimilarValues(t, func(a []int, b int, c ...*uint32) int {
		return find3internal(a, b, 0, c...)
	})
}

func Test_find3_entropy(t *testing.T) {
	goTestEntropic(t, func(a []int, b int, c ...*uint32) int {
		return find3internal(a, b, 0, c...)
	})
}

func Benchmark_find3(b *testing.B) {
	goBenchmarkFind(b, func(a []int, b int) int {
		return find3(a, b)
	})
}
