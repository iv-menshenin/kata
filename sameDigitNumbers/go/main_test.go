package main

import (
	"fmt"
	"testing"
)

func Test_countRepeats(t *testing.T) {
	tests := []struct {
		args int64
		want int
	}{
		{
			args: 999,
			want: 27,
		},
		{
			args: 998,
			want: 26,
		},
		{
			args: 898,
			want: 26,
		},
		{
			args: 798,
			want: 25,
		},
		{
			args: 100,
			want: 18,
		},
		{
			args: 1,
			want: 1,
		},
		{
			args: 100000,
			want: 45,
		},
		{
			args: 236067863,
			want: 74,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.args), func(t *testing.T) {
			if got := countRepeats(tt.args); got != tt.want {
				t.Errorf("countRepeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countRepeatsInRange(t *testing.T) {
	tests := []struct {
		arg1 int64
		arg2 int64
		want int
	}{
		{
			arg1: 13377693,
			arg2: 236067863,
			want: 10,
		},
		{
			arg1: 3217033,
			arg2: 44813321,
			want: 11,
		},
		{
			arg1: 102389076,
			arg2: 166813631,
			want: 1,
		},
		{
			arg1: 7051407,
			arg2: 12126233,
			want: 4,
		},
		{
			arg1: 163257181,
			arg2: 444169533,
			want: 2,
		},
		{
			arg1: 1162530,
			arg2: 1490174,
			want: 0,
		},
		{
			arg1: 57077594,
			arg2: 245289691,
			want: 6,
		},
		{
			arg1: 9013112,
			arg2: 62209491,
			want: 6,
		},
		{
			arg1: 124687807,
			arg2: 367614816,
			want: 2,
		},
		{
			arg1: 394524862,
			arg2: 430707625,
			want: 0,
		},
		{
			arg1: 6917522,
			arg2: 364274799,
			want: 15,
		},
		{
			arg1: 24711101,
			arg2: 295164765,
			want: 9,
		},
		{
			arg1: 29592804,
			arg2: 97983954,
			want: 6,
		},
		{
			arg1: 1867569,
			arg2: 361411176,
			want: 20,
		},
		{
			arg1: 59288264,
			arg2: 211184464,
			want: 5,
		},
		{
			arg1: 24877681,
			arg2: 124369835,
			want: 8,
		},
		{
			arg1: 1,
			arg2: 1000000000000000000,
			want: 162,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d>%d", tt.arg1, tt.arg2), func(t *testing.T) {
			if got := countRepeatsInRange(tt.arg1, tt.arg2); got != tt.want {
				t.Errorf("countRepeatsInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
