package main

import (
	"testing"
)

func Test_find1(t *testing.T) {
	goTestFind(t, func(a []int, b int) int {
		return find1(a, b)
	})
}

func Test_find1_brutForce(t *testing.T) {
	goTestBrutForce(t, func(a []int, b int, c ...*uint32) int {
		return find1(a, b, c...)
	})
}

func Test_find1_similarValues(t *testing.T) {
	goTestSimilarValues(t, func(a []int, b int, c ...*uint32) int {
		return find1(a, b, c...)
	})
}

func Test_find1_entropy(t *testing.T) {
	goTestEntropic(t, func(a []int, b int, c ...*uint32) int {
		return find1(a, b, c...)
	})
}

func Benchmark_find1(b *testing.B) {
	goBenchmarkFind(b, func(a []int, b int) int {
		return find1(a, b)
	})
}
