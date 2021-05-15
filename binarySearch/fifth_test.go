package main

import (
	"testing"
)

func Test_find5(t *testing.T) {
	goTestFind(t, func(a []int, b int) int {
		return find5(a, b)
	})
}

func Test_find5_brutForce(t *testing.T) {
	goTestBrutForce(t, func(a []int, b int, c ...*uint32) int {
		return find5(a, b, c...)
	})
}

func Test_find5_similarValues(t *testing.T) {
	goTestSimilarValues(t, func(a []int, b int, c ...*uint32) int {
		return find5(a, b, c...)
	})
}

func Test_find5_entropy(t *testing.T) {
	goTestEntropic(t, func(a []int, b int, c ...*uint32) int {
		return find5(a, b, c...)
	})
}

func Benchmark_find5(b *testing.B) {
	goBenchmarkFind(b, func(a []int, b int) int {
		return find5(a, b)
	})
}
