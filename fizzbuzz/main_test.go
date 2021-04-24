package main

import (
	"bytes"
	"testing"
)

func Test_isFizz(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "low border case",
			args: args{i: 0},
			want: true,
		},
		{
			name: "one",
			args: args{i: 1},
			want: false,
		},
		{
			name: "fizz",
			args: args{i: 3},
			want: true,
		},
		{
			name: "buzz",
			args: args{i: 5},
			want: false,
		},
		{
			name: "fizzbuzz",
			args: args{i: 30},
			want: true,
		},
		{
			name: "big",
			args: args{i: 30000000000},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFizz(tt.args.i); got != tt.want {
				t.Errorf("isFizz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBuzz(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "low border case",
			args: args{i: 0},
			want: true,
		},
		{
			name: "one",
			args: args{i: 1},
			want: false,
		},
		{
			name: "fizz",
			args: args{i: 3},
			want: false,
		},
		{
			name: "buzz",
			args: args{i: 5},
			want: true,
		},
		{
			name: "fizzbuzz",
			args: args{i: 60},
			want: true,
		},
		{
			name: "big",
			args: args{i: 1000000000},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBuzz(tt.args.i); got != tt.want {
				t.Errorf("isBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fizzBuzz(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "low border case",
			args: args{i: 0},
			want: cFizzBuzz,
		},
		{
			name: "one",
			args: args{i: 1},
			want: "1",
		},
		{
			name: "fizz",
			args: args{i: 66},
			want: cFizz,
		},
		{
			name: "buzz",
			args: args{i: 755},
			want: cBuzz,
		},
		{
			name: "fizzbuzz",
			args: args{i: 750},
			want: cFizzBuzz,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.args.i); got != tt.want {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fizzBuzzRepeat(t *testing.T) {
	type args struct {
		hiRange int
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name:  "zero case",
			args:  args{hiRange: 0},
			wantW: "",
		},
		{
			name:  "from 1 to 3",
			args:  args{hiRange: 3},
			wantW: "1\n2\nFizz\n",
		},
		{
			name:  "from 1 to 5",
			args:  args{hiRange: 5},
			wantW: "1\n2\nFizz\n4\nBuzz\n",
		},
		{
			name:  "from 1 to 15",
			args:  args{hiRange: 15},
			wantW: "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n14\nFizzBuzz\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			fizzBuzzRepeat(tt.args.hiRange, w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("fizzBuzzRepeat() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
