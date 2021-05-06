package main

import (
	"testing"
)

func Test_find4(t *testing.T) {
	goTestFind(t, func(a []int, b int) int {
		return find4(a, b)
	})
}

func Test_find4_brutForce(t *testing.T) {
	goTestBrutForce(t, func(a []int, b int, c ...*uint32) int {
		return find4internal(a, b, 0, c...)
	})
}

func Test_find4_similarValues(t *testing.T) {
	goTestSimilarValues(t, func(a []int, b int, c ...*uint32) int {
		return find4internal(a, b, 0, c...)
	})
}

func Test_find4_entropy(t *testing.T) {
	goTestEntropic(t, func(a []int, b int, c ...*uint32) int {
		return find4internal(a, b, 0, c...)
	})
}

func Benchmark_find4(b *testing.B) {
	goBenchmarkFind(b, func(a []int, b int) int {
		return find4(a, b)
	})
}
