package loginCounter

import "testing"

func Benchmark_counter_mapBucketLogger(b *testing.B) {
	/*
		cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz
		Benchmark_counter_mapBucketLogger
		Benchmark_counter_mapBucketLogger-8     	  643489	     44557 ns/op	   66298 B/op	      36 allocs/op
	*/
	var c = NewMapBucketLogger()
	benchCounter(b, c)
}

func Test_mapBucketLogger_Performance(t *testing.T) {
	var c = NewMapBucketLogger()
	testCounterPerformance(t, c)
}
