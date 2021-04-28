package main

import (
	"math"
	"sort"
)

type (
	caseT struct {
		nn float64
		cp string
	}
	caseA []caseT
)

func (a caseA) Len() int {
	return len(a)
}

func (a caseA) Less(i, j int) bool {
	return math.Abs(a[i].nn) < math.Abs(a[j].nn)
}

func (a caseA) Swap(i, j int) {
	a[i].nn, a[j].nn = a[j].nn, a[i].nn
	a[i].cp, a[j].cp = a[j].cp, a[i].cp
}

func testComplexity(items, ticks int) string {
	var aa caseA
	aa = append(aa, caseT{
		nn: 1.0 - float64(ticks),
		cp: "O(1)",
	})
	aa = append(aa, caseT{
		nn: math.Log(float64(items)) - float64(ticks),
		cp: "O(log n)",
	})
	aa = append(aa, caseT{
		nn: float64(items - ticks),
		cp: "O(n)",
	})
	aa = append(aa, caseT{
		nn: float64(items)*math.Log(float64(items)) - float64(ticks),
		cp: "O(n * log n)",
	})
	aa = append(aa, caseT{
		nn: float64(items ^ 2 - ticks),
		cp: "O(n^2)",
	})
	// O(!n) unsupported
	sort.Sort(aa)
	return aa[0].cp
}
