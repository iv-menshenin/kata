package wrapper

import (
	"reflect"
	"strings"
	"testing"
)

func Test_Wrap(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name   string
		input  string
		colSz  int
		output []string
	}
	tests := []testCase{
		{
			name:  "empty",
			colSz: 1,
		},
		{
			name:   "short_word",
			input:  "test",
			colSz:  5,
			output: []string{"test"},
		},
		{
			name:   "sort_wrapped",
			input:  "test",
			colSz:  2,
			output: []string{"te", "st"},
		},
		{
			name:   "split_by_space",
			input:  "test some",
			colSz:  5,
			output: []string{"test", "some"},
		},
		{
			name:   "split_by_space",
			input:  "we didn't have much time",
			colSz:  9,
			output: []string{"we didn't", "have much", "time"},
		},
		{
			name:   "yet_another_test",
			input:  "yet another test",
			colSz:  3,
			output: []string{"yet", "ano", "the", "r", "tes", "t"},
		},
	}
	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := Wrap(test.input, test.colSz)
			if !reflect.DeepEqual(got, test.output) {
				t.Errorf("matching error\nwant: %v\n got: %v", test.output, strings.Join(got, ","))
			}
		})
	}
}
