package loginCounter

import "testing"

func Benchmark_counter_arrayBucketLogger(b *testing.B) {
	/*
		cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz
		Benchmark_counter_arrayBucketLogger
		Benchmark_counter_arrayBucketLogger-8   	  383007	    197112 ns/op	   47239 B/op	      22 allocs/op
	*/
	var c = NewArrayBucketLogger()
	benchCounter(b, c)
}

func Test_arrayBucketLogger_Performance(t *testing.T) {
	var c = NewArrayBucketLogger()
	testCounterPerformance(t, c)
}
