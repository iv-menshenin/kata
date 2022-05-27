package collapseNumeric

import "testing"

type collapseNumericFn func([]int) string

func collapseNumericSequenceTest(t *testing.T, fn collapseNumericFn) {
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

func Test_collapseNumericSequenceOOP(t *testing.T) {
	collapseNumericSequenceTest(t, collapseNumericSequenceOOP)
}

func Test_collapseNumericSequenceIter(t *testing.T) {
	collapseNumericSequenceTest(t, collapseNumericSequenceIter)
}

func Test_collapseNumericSequenceOptimistic(t *testing.T) {
	collapseNumericSequenceTest(t, collapseNumericSequenceOptimistic)
}

func benchCollapseNumericSequence(b *testing.B, data []int, fn collapseNumericFn) {
	for i := 1; i < b.N; i++ {
		_ = fn(data)
	}
}

func Benchmark_collapseNumericSequence(t *testing.B) {
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
	t.Run("OOP", func(b *testing.B) {
		benchCollapseNumericSequence(b, data, collapseNumericSequenceOOP)
	})
	t.Run("Iter", func(b *testing.B) {
		benchCollapseNumericSequence(b, data, collapseNumericSequenceIter)
	})
	t.Run("Optimistic", func(b *testing.B) {
		benchCollapseNumericSequence(b, data, collapseNumericSequenceOptimistic)
	})
	/*
		cpu: Intel(R) Core(TM) i7-9700F CPU @ 3.00GHz
		Benchmark_collapseNumericSequence
		Benchmark_collapseNumericSequence/OOP
		Benchmark_collapseNumericSequence/OOP-8         	  428552	      2770 ns/op	     640 B/op	      17 allocs/op
		Benchmark_collapseNumericSequence/Iter
		Benchmark_collapseNumericSequence/Iter-8        	  407670	      2937 ns/op	     720 B/op	      29 allocs/op
		Benchmark_collapseNumericSequence/Optimistic
		Benchmark_collapseNumericSequence/Optimistic-8  	  698198	      1665 ns/op	     167 B/op	       2 allocs/op
	*/
}
