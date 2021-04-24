// run it with --benchmem
package main

import (
	"testing"
)

func Benchmark_awesomeFuncWithStruct(b *testing.B) {
	var i = 0
	for {
		if i >= b.N {
			break
		}
		a, e := awesomeFuncWithStruct(i)
		if e != nil && a != nil {
			println("foo")
		}
		i++
	}
}

func Benchmark_awesomeFuncWithSlice(b *testing.B) {
	var i = 0
	for {
		if i >= b.N {
			break
		}
		a, e := awesomeFuncWithSlice(i)
		if e != nil && len(a) > 0 {
			println("foo")
		}
		i++
	}
}

func Benchmark_awesomeFuncWithMake(b *testing.B) {
	var i = 0
	for {
		if i >= b.N {
			break
		}
		a, e := awesomeFuncWithMake(i)
		if e != nil && len(a) > 0 {
			println("foo")
		}
		i++
	}
}
