package scrollupNumbers

import "testing"

type scrollupNumbersFn func([]int) string

func testScrollupNumbers(t *testing.T, fn scrollupNumbersFn) {
	tests := []struct {
		name string
		arg  []int
		want string
	}{
		{
			name: "nil",
			arg:  nil,
			want: "",
		},
		{
			name: "once",
			arg:  []int{4},
			want: "4",
		},
		{
			name: "one_tuple",
			arg:  []int{2, 1, 0, 4, 3},
			want: "0-4",
		},
		{
			name: "mixed",
			arg:  []int{12, 0, 2, 11, 3, 4, 5, 6, 13},
			want: "0,2-6,11-13",
		},
		{
			name: "mixed_with_appendix",
			arg:  []int{12, 2, 11, 3, 44, 4, 5, 6, 13},
			want: "2-6,11-13,44",
		},
	}
	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			if got := fn(test.arg); got != test.want {
				t.Errorf("got: '%s', want: '%s'", got, test.want)
			}
		})
	}
}

func Test_scrollupNumbersSerial(t *testing.T) {
	testScrollupNumbers(t, scrollupNumbersSerial)
}

func Test_scrollNumbersBoilerPrint(t *testing.T) {
	testScrollupNumbers(t, scrollNumbersBoilerPrint)
}

func Test_scrollNumbersSimpled(t *testing.T) {
	testScrollupNumbers(t, scrollNumbersSimpled)
}

func Test_scrollNumbersOptimistic(t *testing.T) {
	testScrollupNumbers(t, scrollNumbersOptimistic)
}

func benchScrollupNumbers(b *testing.B, data []int, fn scrollupNumbersFn) {
	for i := 1; i < b.N; i++ {
		_ = fn(data)
	}
}

func Benchmark_scrollupNumbers(t *testing.B) {
	var data []int
	for i := 0; i < 100; i++ {
		if i%13 == 0 {
			continue
		}
		if i%27 == 0 {
			continue
		}
		data = append(data, i)
	}
	t.ResetTimer()
	t.Run("scrollupNumbersSerial", func(b *testing.B) {
		benchScrollupNumbers(b, data, scrollupNumbersSerial)
	})
	t.Run("scrollNumbersBoilerPrint", func(b *testing.B) {
		benchScrollupNumbers(b, data, scrollNumbersBoilerPrint)
	})
	t.Run("scrollNumbersSimpled", func(b *testing.B) {
		benchScrollupNumbers(b, data, scrollNumbersSimpled)
	})
	t.Run("scrollNumbersOptimistic", func(b *testing.B) {
		benchScrollupNumbers(b, data, scrollNumbersOptimistic)
	})
	/*
		   cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz
		   	Benchmark_scrollupNumbers
		   	Benchmark_scrollupNumbers/scrollupNumbersSerial
			Benchmark_scrollupNumbers/scrollupNumbersSerial-8         	  379347	      3133 ns/op	     640 B/op	      17 allocs/op
			Benchmark_scrollupNumbers/scrollNumbersBoilerPrint
			Benchmark_scrollupNumbers/scrollNumbersBoilerPrint-8      	  357234	      3389 ns/op	     720 B/op	      29 allocs/op
			Benchmark_scrollupNumbers/scrollNumbersSimpled
			Benchmark_scrollupNumbers/scrollNumbersSimpled-8          	  358098	      3347 ns/op	     720 B/op	      29 allocs/op
			Benchmark_scrollupNumbers/scrollNumbersOptimistic
			Benchmark_scrollupNumbers/scrollNumbersOptimistic-8       	  602614	      2005 ns/op	     167 B/op	       2 allocs/op
	*/
}
