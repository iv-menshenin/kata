package loginCounter

import "testing"

func Benchmark_counter_arrayBucketLogger_Login(b *testing.B) {
	/*
		cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz
		Benchmark_counter_arrayBucketLogger_Login-8   	 1000000	      1339 ns/op	     277 B/op	       0 allocs/op
	*/
	var c = NewArrayBucketLogger()
	benchCounterLogin(b, c)
}

func Benchmark_counter_arrayBucketLogger_Count(b *testing.B) {
	/*
		cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz
		Benchmark_counter_arrayBucketLogger_Count-8   	  347372	     15536 ns/op	      82 B/op	       0 allocs/op
	*/
	var c = NewArrayBucketLogger()
	benchCounterCount(b, c)
}

func Test_arrayBucketLogger_Performance(t *testing.T) {
	var c = NewArrayBucketLogger()
	testCounterPerformance(t, c)
}

func Test_arrayBucketLogger_Functional(t *testing.T) {
	var c = NewArrayBucketLogger()
	testCounterCountFunc(t, c)
}
