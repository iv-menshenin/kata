package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test_find1(t *testing.T) {
	type args struct {
		a []int
		f int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "midden 1",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17, 21},
				f: 7,
			},
			want: 5,
		},
		{
			name: "midden 2",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17, 21},
				f: 17,
			},
			want: 8,
		},
		{
			name: "midden 3",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17, 21},
				f: 2,
			},
			want: 2,
		},
		{
			name: "first 1",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17, 21},
				f: 0,
			},
			want: 0,
		},
		{
			name: "first 2",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17},
				f: 0,
			},
			want: 0,
		},
		{
			name: "first 3",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13},
				f: 0,
			},
			want: 0,
		},
		{
			name: "last 1",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17, 21},
				f: 21,
			},
			want: 9,
		},
		{
			name: "last 2",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13, 17},
				f: 17,
			},
			want: 8,
		},
		{
			name: "last 3",
			args: args{
				a: []int{0, 1, 2, 3, 5, 7, 11, 13},
				f: 13,
			},
			want: 7,
		},
		{
			name: "nil array",
			args: args{
				a: nil,
				f: 13,
			},
			want: -1,
		},
		{
			name: "empty array",
			args: args{
				a: []int{},
				f: 13,
			},
			want: -1,
		},
		{
			name: "short array 1",
			args: args{
				a: []int{1},
				f: 1,
			},
			want: 0,
		},
		{
			name: "short array 2",
			args: args{
				a: []int{1, 2},
				f: 1,
			},
			want: 0,
		},
		{
			name: "short array 3",
			args: args{
				a: []int{1, 2},
				f: 2,
			},
			want: 1,
		},
		{
			name: "short array 4",
			args: args{
				a: []int{1, 2, 3},
				f: 3,
			},
			want: 2,
		},
		{
			name: "not found 1",
			args: args{
				a: []int{1, 2, 3},
				f: 6,
			},
			want: -1,
		},
		{
			name: "not found 2",
			args: args{
				a: []int{1, 2, 4, 5},
				f: 3,
			},
			want: -1,
		},
		{
			name: "not found 3",
			args: args{
				a: []int{1},
				f: 6,
			},
			want: -1,
		},
		{
			name: "not found 4",
			args: args{
				a: []int{1, 10, 100, 1000},
				f: 60,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find1(tt.args.a, tt.args.f); got != tt.want {
				t.Errorf("find1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_find1_brutForce(t *testing.T) {
	var (
		a     = make([]int, 0, 1000)
		tests = make(map[int]int)
	)
	for i := 0; i < cap(a); i++ {
		a = append(a, i*3)
	}
	for i := 0; i < 10; i++ {
		idx := rand.Intn(cap(a))
		tests[i] = idx
	}
	for k, v := range tests {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			var iterCount uint32
			if got := find1(a, a[v], &iterCount); got != v {
				t.Errorf("find1() = %v, want %v", got, v)
			}
			println(fmt.Sprintf("#%d: %d", k, iterCount))
		})
	}
}

func Benchmark_find1(b *testing.B) {
	var (
		a     = make([]int, 0, 1000)
		tests = make(map[int]int)
	)
	for i := 0; i < cap(a); i++ {
		a = append(a, i)
	}
	tests[1] = 600
	tests[2] = 1
	tests[3] = 998
	tests[4] = 124
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			if got := find1(a, v); got != v {
				b.Errorf("find1() = %v, want %v", got, v)
			}
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
