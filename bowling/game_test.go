package bowling

import "testing"

func Test_game_rollZero(t *testing.T) {
	var g = gameScores{}
	for nn := 0; nn < 20; nn++ {
		g.Roll(0)
	}
	if score := g.Score(); score != 0 {
		t.Fatal("expected 0, got", score)
	}
}

func Test_game_rollOne(t *testing.T) {
	var g = gameScores{}
	for nn := 0; nn < 20; nn++ {
		g.Roll(1)
	}
	if score := g.Score(); score != 20 {
		t.Fatal("expected 20, got", score)
	}
}

func Test_game_spareOne(t *testing.T) {
	var g = gameScores{}
	g.Roll(5)
	g.Roll(5)
	g.Roll(3)
	if score := g.Score(); score != 16 {
		t.Fatal("expected 16, got", score)
	}

	g = gameScores{}
	g.Roll(5)
	g.Roll(5)
	g.Roll(3)
	g.Roll(0)
	g.Roll(0)
	if score := g.Score(); score != 16 {
		t.Fatal("expected 16, got", score)
	}
}

func Test_game_strikeOne(t *testing.T) {
	var g = gameScores{}
	g.Roll(10)
	g.Roll(3)
	g.Roll(4)
	if score := g.Score(); score != 24 {
		t.Fatal("expected 24, got", score)
	}

	g = gameScores{}
	g.Roll(10)
	g.Roll(3)
	g.Roll(4)
	g.Roll(0)
	g.Roll(0)
	if score := g.Score(); score != 24 {
		t.Fatal("expected 24, got", score)
	}
}

func Test_game_best(t *testing.T) {
	var g = gameScores{}
	for nn := 0; nn < 12; nn++ {
		g.Roll(10)
	}
	if score := g.Score(); score != 300 {
		t.Fatal("expected 300, got", score)
	}
}

func Test_frame_isSpare(t *testing.T) {
	tests := []struct {
		name  string
		score []int
		want  bool
	}{
		{name: "0_10", score: []int{0, 10}, want: true},
		{name: "1_8", score: []int{1, 8}, want: false},
		{name: "1_9", score: []int{1, 9}, want: true},
		{name: "5_5", score: []int{5, 5}, want: true},
		{name: "_10", score: []int{10}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := gameFrame{score: tt.score}
			if got := f.isSpare(); got != tt.want {
				t.Errorf("isSpare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_frame_isStrike(t *testing.T) {
	tests := []struct {
		name  string
		score []int
		want  bool
	}{
		{name: "0_10", score: []int{0, 10}, want: false},
		{name: "1_8", score: []int{1, 8}, want: false},
		{name: "1_9", score: []int{1, 9}, want: false},
		{name: "5_5", score: []int{5, 5}, want: false},
		{name: "_10", score: []int{10}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := gameFrame{score: tt.score}
			if got := f.isStrike(); got != tt.want {
				t.Errorf("isStrike() = %v, want %v", got, tt.want)
			}
		})
	}
}
