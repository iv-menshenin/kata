package binaryVector

import "testing"

func Test_countSerials(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name   string
		vector []int8
		need   int
	}{
		{
			name:   "started_with_1",
			vector: []int8{1, 0, 0, 0, 0},
			need:   1,
		},
		{
			name:   "ended_with_1",
			vector: []int8{0, 0, 0, 0, 0, 0, 1},
			need:   1,
		},
		{
			name:   "simple_1",
			vector: []int8{0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 0},
			need:   3,
		},
		{
			name:   "simple_2",
			vector: []int8{0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 1, 0},
			need:   3,
		},
		{
			name:   "five_ones",
			vector: []int8{1, 1, 1, 1, 1},
			need:   5,
		},
		{
			name:   "five_zeroes",
			vector: []int8{0, 0, 0, 0, 0},
			need:   0,
		},
		{
			name:   "nil",
			vector: nil,
			need:   0,
		},
	}
	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := countSerials(test.vector); test.need != got {
				t.Errorf("need: %v, got: %v", test.need, got)
			}
		})
	}
}
