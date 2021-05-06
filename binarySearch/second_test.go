package main

import (
	"testing"
)

func Test_find2(t *testing.T) {
	goTestFind(t, func(a []int, b int) int {
		return find2(a, b)
	})
}

func Test_find2_brutForce(t *testing.T) {
	goTestBrutForce(t, func(a []int, b int, c ...*uint32) int {
		return find2(a, b, c...)
	})
}

func Test_find2_similarValues(t *testing.T) {
	goTestSimilarValues(t, func(a []int, b int, c ...*uint32) int {
		return find2(a, b, c...)
	})
}

func Test_find2_entropy(t *testing.T) {
	goTestEntropic(t, func(a []int, b int, c ...*uint32) int {
		return find2(a, b, c...)
	})
}

func Benchmark_find2(b *testing.B) {
	goBenchmarkFind(b, func(a []int, b int) int {
		return find2(a, b)
	})
}
