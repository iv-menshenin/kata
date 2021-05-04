package main

import (
	"math/rand"
	"testing"
	"time"
)

func Benchmark_some1_MarshalAMQP(b *testing.B) {
	var a []some1
	for nn := 0; nn < b.N; nn++ {
		a = append(a, some1{RequestID: rand.Int63()})
	}
	b.ResetTimer()
	for _, s := range a {
		_, _ = s.MarshalAMQP()
	}
}

func Benchmark_some2_MarshalAMQP(b *testing.B) {
	var a []some2
	for nn := 0; nn < b.N; nn++ {
		a = append(a, some2{RequestID: rand.Int63()})
	}
	b.ResetTimer()
	for _, s := range a {
		_, _ = s.MarshalAMQP()
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
