package loginCounter

import "testing"

func Benchmark_counter_mapBucketLogger_Login(b *testing.B) {
	/*
		cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz
		Benchmark_counter_mapBucketLogger_Login-8     	 1000000	      1387 ns/op	     153 B/op	       0 allocs/op
	*/
	var c = NewMapBucketLogger()
	benchCounterLogin(b, c)
}

func Benchmark_counter_mapBucketLogger_Count(b *testing.B) {
	/*
		cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz
		Benchmark_counter_mapBucketLogger_Count-8     	 1871526	       672.8 ns/op	       0 B/op	       0 allocs/op
	*/
	var c = NewMapBucketLogger()
	benchCounterCount(b, c)
}

func Test_mapBucketLogger_Performance(t *testing.T) {
	var c = NewMapBucketLogger()
	testCounterPerformance(t, c)
}

func Test_mapBucketLogger_Functional(t *testing.T) {
	var c = NewMapBucketLogger()
	testCounterCountFunc(t, c)
}
