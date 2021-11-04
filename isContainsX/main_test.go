package isContainsX

import (
	"math/rand"
	"testing"
	"time"
)

func Test_byArraySlow(t *testing.T) {
	testHelper(t, byArraySlow)
}

func Test_byIntArr(t *testing.T) {
	testHelper(t, byIntArr)
}

func Test_byMap(t *testing.T) {
	testHelper(t, byMap)
}

func Test_byMap2(t *testing.T) {
	testHelper(t, byMap2)
}

func testHelper(t *testing.T, f func([]int) bool) {
	t.Helper()
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "border-1",
			args: args{arr: []int{1, 0}},
			want: false,
		},
		{
			name: "border-2",
			args: args{arr: []int{3, 1}},
			want: false,
		},
		{
			name: "sorted-4",
			args: args{arr: []int{1, 3, 6, 10}},
			want: true,
		},
		{
			name: "unsorted-4",
			args: args{arr: []int{2, 7, 12, 6}},
			want: true,
		},
		{
			name: "sorted-13",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
			want: true,
		},
		{
			name: "unsorted-simple",
			args: args{arr: []int{1, 3, 5, 8, 13, 18, 31, 49, 90, 139}},
			want: false,
		},
		{
			name: "simple",
			args: args{arr: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.args.arr); got != tt.want {
				t.Errorf("f() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchFunc(b *testing.B, f func([]int) bool, nnn int) {
	b.Helper()
	rand.Seed(time.Now().UnixNano())
	var arr = make([]int, 0)
	for nn := 0; nn < nnn; nn++ {
		arr = append(arr, rand.Intn(1000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(arr)
	}
}

func Benchmark_byArraySlow(b *testing.B) {
	BenchFunc(b, byArraySlow, 500)
}

func Benchmark_byMap(b *testing.B) {
	BenchFunc(b, byMap, 500)
}

func Benchmark_byIntArr(b *testing.B) {
	BenchFunc(b, byIntArr, 500)
}

func Benchmark_byMap2(b *testing.B) {
	BenchFunc(b, byMap2, 500)
}
