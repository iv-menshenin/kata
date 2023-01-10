package topfreq

import (
	"reflect"
	"testing"
)

func Test_topFreq(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name string
		arg1 []int64
		arg2 int
		want []int64
	}{
		{
			name: "simple_one",
			arg1: []int64{12},
			arg2: 1,
			want: []int64{12},
		},
		{
			name: "cut_results",
			arg1: []int64{12, 5},
			arg2: 1,
			want: []int64{12},
		},
		{
			name: "group_all",
			arg1: []int64{12, 12, 5},
			arg2: 1,
			want: []int64{12},
		},
		{
			name: "sort_all",
			arg1: []int64{12, 5, 12},
			arg2: 3,
			want: []int64{12, 5},
		},
		{
			name: "sort_by_freq",
			arg1: []int64{1, 2, 1, 1, 3, 2},
			arg2: 3,
			want: []int64{1, 2, 3},
		},
		{
			name: "empty",
			arg1: []int64{},
			arg2: 1,
			want: nil,
		},
		{
			name: "issue-0001",
			arg1: []int64{1, 2, 1, 1, 3, 2, 54, 66, 10, 11, 11, 13, 14, 1},
			arg2: 3,
			want: []int64{1, 2, 11},
		},
	}
	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := getTopFrequent(test.arg1, test.arg2)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("matching error\n got: %v\nwant: %v", got, test.want)
			}
		})
	}
}
