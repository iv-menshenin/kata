package loginCounter

import "testing"

func Benchmark_counter_mapBucketLogger(b *testing.B) {
	/*
		cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz
		Benchmark_counter_mapBucketLogger
		Benchmark_counter_mapBucketLogger-8   	  779850	     13804 ns/op	    9050 B/op	       3 allocs/op
	*/
	var c = NewMapBucketLogger()
	benchCounter(b, c)
}

func Test_mapBucketLogger_Performance(t *testing.T) {
	var c = NewMapBucketLogger()
	testCounterPerformance(t, c)
}
