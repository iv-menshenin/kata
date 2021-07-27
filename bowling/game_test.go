package bowling

import "testing"

func Test_game_rollZero(t *testing.T) {
	var g = Game{}
	for nn := 0; nn < 20; nn++ {
		g.Roll(0)
	}
	if score := g.Score(); score != 0 {
		t.Fatal("expected 0, got", score)
	}
}

func Test_game_rollOne(t *testing.T) {
	var g = Game{}
	for nn := 0; nn < 20; nn++ {
		g.Roll(1)
	}
	if score := g.Score(); score != 20 {
		t.Fatal("expected 20, got", score)
	}
}

func Test_game_spareOne(t *testing.T) {
	var g = Game{}
	g.Roll(5)
	g.Roll(5)
	g.Roll(3)
	if score := g.Score(); score != 16 {
		t.Fatal("expected 16, got", score)
	}

	g = Game{}
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
	var g = Game{}
	g.Roll(10)
	g.Roll(3)
	g.Roll(4)
	if score := g.Score(); score != 24 {
		t.Fatal("expected 24, got", score)
	}

	g = Game{}
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
	var g = Game{}
	for nn := 0; nn < 12; nn++ {
		g.Roll(10)
	}
	if score := g.Score(); score != 300 {
		t.Fatal("expected 300, got", score)
	}
}
