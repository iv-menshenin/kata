package main

import (
	"bytes"
	"testing"
)

func Test_fizzBuzzRepeatChan(t *testing.T) {
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
			fizzBuzzRepeatChan(tt.args.hiRange, w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("fizzBuzzRepeatChan() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
