package bowling

import "testing"

func Test_Bowling(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name  string
		rolls []int
		score int
	}
	tests := []testCase{
		{
			name:  "all_zero",
			rolls: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			score: 0,
		},
		{
			name:  "some_strange_1",
			rolls: []int{0, 10, 0, 0, 1, 0, 0, 0, 0, 0},
			score: 11,
		},
		{
			name:  "some_strange_2",
			rolls: []int{0, 10, 0, 1, 0, 0, 0, 0, 0},
			score: 11,
		},
		{
			name:  "some_strange_3",
			rolls: []int{10, 0, 1, 0, 0, 0, 0, 0},
			score: 12,
		},
		{
			name:  "roll_once",
			rolls: []int{4},
			score: 4,
		},
		{
			name:  "roll_zero",
			rolls: []int{},
			score: 0,
		},
		{
			name:  "double_roll",
			rolls: []int{2, 1},
			score: 3,
		},
		{
			name:  "spare",
			rolls: []int{4, 6, 2},
			score: 14,
		},
		{
			name:  "double_spare",
			rolls: []int{4, 6, 2, 8, 1},
			score: 24,
		},
		{
			name:  "strike",
			rolls: []int{10, 8, 1},
			score: 28,
		},
		{
			name:  "spare_n_strike",
			rolls: []int{6, 4, 10, 4, 1},
			score: 40,
		},
		{
			name:  "strike_n_nostrike",
			rolls: []int{10, 1, 2, 5},
			score: 21,
		},
		{
			name:  "double_strike",
			rolls: []int{10, 10, 2, 5},
			score: 46,
		},
		{
			name:  "strike_n_spare",
			rolls: []int{10, 6, 4, 2},
			score: 34,
		},
		{
			name:  "max_score",
			rolls: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			score: 300,
		},
	}
	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			game := New()
			for _, roll := range test.rolls {
				game.Roll(roll)
			}
			got := game.Score()
			if got != test.score {
				t.Errorf("matching error\ngot:  %d\nwant: %d", got, test.score)
			}
		})
	}
}
