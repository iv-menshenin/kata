package factor

import (
	"reflect"
	"testing"
)

func Test_Generate(t *testing.T) {
	t.Parallel()
	type testCase = struct {
		name string
		arg  int
		want []int
	}
	var tests = []testCase{
		{
			name: "one",
			arg:  1,
			want: nil,
		},
		{
			name: "two",
			arg:  2,
			want: []int{2},
		},
		{
			name: "three",
			arg:  3,
			want: []int{3},
		},
		{
			name: "four",
			arg:  4,
			want: []int{2, 2},
		},
		{
			name: "nine",
			arg:  9,
			want: []int{3, 3},
		},
		{
			name: "twenty_six",
			arg:  26,
			want: []int{2, 13},
		},
	}
	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			got := Generate(test.arg)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("matching error\nwant: %v\n got: %v", test.want, got)
			}
		})
	}
}
