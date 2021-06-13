// run it with --benchmem
package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func Test_normalizeViewActions(t *testing.T) {
	type args struct {
		sessions viewSessions
	}
	tests := []struct {
		name string
		args args
		want viewEvents
	}{
		{
			name: "empty",
			args: args{},
			want: viewEvents{},
		},
		{
			name: "3 case",
			args: args{sessions: []viewSession{
				{start: 2, end: 9},
				{start: 15, end: 21},
				{start: 3, end: 12},
			}},
			want: viewEvents{
				{isIncrease: true, timePoint: 2},
				{isIncrease: true, timePoint: 3},
				{isIncrease: false, timePoint: 9},
				{isIncrease: false, timePoint: 12},
				{isIncrease: true, timePoint: 15},
				{isIncrease: false, timePoint: 21},
			},
		},
		{
			name: "cross-case 1",
			args: args{sessions: []viewSession{
				{start: 18, end: 78},
				{start: 24, end: 34},
				{start: 16, end: 90},
				{start: 2, end: 56},
				{start: 56, end: 88},
			}},
			want: viewEvents{
				{isIncrease: true, timePoint: 2},
				{isIncrease: true, timePoint: 16},
				{isIncrease: true, timePoint: 18},
				{isIncrease: true, timePoint: 24},
				{isIncrease: false, timePoint: 34},
				{isIncrease: false, timePoint: 56},
				{isIncrease: true, timePoint: 56},
				{isIncrease: false, timePoint: 78},
				{isIncrease: false, timePoint: 88},
				{isIncrease: false, timePoint: 90},
			},
		},
		{
			name: "cross-case 2",
			args: args{sessions: []viewSession{
				{start: 18, end: 78},
				{start: 24, end: 34},
				{start: 56, end: 88},
				{start: 16, end: 90},
				{start: 2, end: 56},
			}},
			want: viewEvents{
				{isIncrease: true, timePoint: 2},
				{isIncrease: true, timePoint: 16},
				{isIncrease: true, timePoint: 18},
				{isIncrease: true, timePoint: 24},
				{isIncrease: false, timePoint: 34},
				{isIncrease: false, timePoint: 56},
				{isIncrease: true, timePoint: 56},
				{isIncrease: false, timePoint: 78},
				{isIncrease: false, timePoint: 88},
				{isIncrease: false, timePoint: 90},
			},
		},
		{
			name: "one after another",
			args: args{sessions: []viewSession{
				{start: 5, end: 8},
				{start: 12, end: 90},
				{start: 8, end: 12},
				{start: 1, end: 5},
			}},
			want: viewEvents{
				{isIncrease: true, timePoint: 1},
				{isIncrease: false, timePoint: 5},
				{isIncrease: true, timePoint: 5},
				{isIncrease: false, timePoint: 8},
				{isIncrease: true, timePoint: 8},
				{isIncrease: false, timePoint: 12},
				{isIncrease: true, timePoint: 12},
				{isIncrease: false, timePoint: 90},
			},
		},
		{
			name: "all crossed",
			args: args{sessions: []viewSession{
				{start: 123, end: 677},
				{start: 5, end: 8},
				{start: 5, end: 8},
				{start: 8, end: 12},
				{start: 8, end: 12},
			}},
			want: viewEvents{
				{isIncrease: true, timePoint: 5},
				{isIncrease: true, timePoint: 5},
				{isIncrease: false, timePoint: 8},
				{isIncrease: false, timePoint: 8},
				{isIncrease: true, timePoint: 8},
				{isIncrease: true, timePoint: 8},
				{isIncrease: false, timePoint: 12},
				{isIncrease: false, timePoint: 12},
				{isIncrease: true, timePoint: 123},
				{isIncrease: false, timePoint: 677},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeViewActions(tt.args.sessions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalizeViewActions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateMaxViewers1(t *testing.T) {
	type args struct {
		events viewEvents
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero",
			args: args{events: viewEvents{}},
			want: 0,
		},
		{
			name: "6 users case 1",
			args: args{events: viewEvents{
				{isIncrease: true, timePoint: 190},
				{isIncrease: true, timePoint: 200},
				{isIncrease: true, timePoint: 212},
				{isIncrease: true, timePoint: 224},
				{isIncrease: false, timePoint: 237},
				{isIncrease: true, timePoint: 315},
				{isIncrease: false, timePoint: 334},
				{isIncrease: false, timePoint: 456},
				{isIncrease: true, timePoint: 890},
				{isIncrease: false, timePoint: 1024},
				{isIncrease: false, timePoint: 1056},
				{isIncrease: false, timePoint: 1099},
			}},
			want: 4,
		},
		{
			name: "6 users case 2 crossed",
			args: args{events: viewEvents{
				{isIncrease: true, timePoint: 190},
				{isIncrease: false, timePoint: 200},
				{isIncrease: true, timePoint: 212},
				{isIncrease: false, timePoint: 224},
				{isIncrease: true, timePoint: 237},
				{isIncrease: false, timePoint: 315},
				{isIncrease: false, timePoint: 456},
				{isIncrease: true, timePoint: 456},
				{isIncrease: true, timePoint: 890},
				{isIncrease: false, timePoint: 1024},
				{isIncrease: true, timePoint: 1056},
				{isIncrease: false, timePoint: 1099},
			}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMaxViewers1(tt.args.events); got != tt.want {
				t.Errorf("calculateMaxViewers1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_calculateMax1Viewers(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	var ses = make(viewSessions, b.N, b.N)
	for i := 0; i < len(ses); i++ {
		ses[i].start = rand.Int63n(1000000)
		ses[i].end = ses[i].start + rand.Int63n(1000000)
	}
	b.ResetTimer()
	var evt = normalizeViewActions(ses)
	calculateMaxViewers1(evt)
}

func Benchmark_calculateMax2Viewers(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	var ses = make(viewSessions, b.N, b.N)
	for i := 0; i < len(ses); i++ {
		ses[i].start = rand.Int63n(1000000)
		ses[i].end = ses[i].start + rand.Int63n(1000000)
	}
	b.ResetTimer()
	calculateMaxViewers2(ses)
}

func Test_calculateMaxViewers2(t *testing.T) {
	type args struct {
		sessions viewSessions
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero",
			args: args{},
			want: 0,
		},
		{
			name: "9 max",
			args: args{sessions: viewSessions{
				{start: 17, end: 56},
				{start: 23, end: 78},
				{start: 1, end: 60},
				{start: 12, end: 34},
				{start: 14, end: 56},
				{start: 2, end: 45},
				{start: 60, end: 80},
				{start: 24, end: 80},
				{start: 24, end: 80},
				{start: 1, end: 40},
				{start: 1, end: 12},
			}},
			want: 9,
		},
		{
			name: "4 max",
			args: args{sessions: viewSessions{
				{start: 2, end: 17},
				{start: 12, end: 17},
				{start: 24, end: 80},
				{start: 17, end: 56},
				{start: 1, end: 17},
				{start: 24, end: 80},
				{start: 1, end: 17},
				{start: 23, end: 78},
				{start: 60, end: 80},
				{start: 1, end: 12},
			}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMaxViewers2(tt.args.sessions); got != tt.want {
				t.Errorf("calculateMaxViewers2() = %v, want %v", got, tt.want)
			}
		})
	}
}
