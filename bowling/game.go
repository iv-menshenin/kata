package bowling

type (
	bonus int
	frame struct {
		score      []int
		over       []int
		frameBonus bonus
	}
	Game struct {
		frames []frame
	}
)

const (
	regular bonus = iota
	spare
	strike
)

func scoreFrame(f []int) (result int) {
	for _, s := range f {
		result += s
	}
	return result
}

func (g *Game) Roll(score int) {
	var f frame
	lastFrame := len(g.frames) - 1
	if lastFrame < 0 || len(g.frames[lastFrame].score) == 2 || scoreFrame(g.frames[lastFrame].score) > 9 {
		if lastFrame == 9 {
			f = g.frames[lastFrame]
			f.over = append(f.over, score)
			g.frames[lastFrame] = f
			return
		}
		f = frame{
			score:      make([]int, 0, 3),
			frameBonus: regular,
		}
		g.frames = append(g.frames, f)
		lastFrame = len(g.frames) - 1
	} else {
		f = g.frames[lastFrame]
	}
	f.score = append(f.score, score)
	if f.frameBonus == regular && scoreFrame(f.score) == 10 && len(f.score) == 2 {
		f.frameBonus = spare
	}
	if f.frameBonus == regular && scoreFrame(f.score) == 10 && len(f.score) == 1 {
		f.frameBonus = strike
	}
	g.frames[lastFrame] = f
}

func (g *Game) Score() (result int) {
	var bonus []int
	for _, f := range g.frames {
		for _, s := range f.score {
			result += s
			for i := range bonus {
				if bonus[i] > 0 {
					bonus[i]--
					result += s
				}
			}
		}
		if f.frameBonus == spare {
			bonus = append(bonus, 1)
		}
		if f.frameBonus == strike {
			bonus = append(bonus, 2)
		}
		for _, s := range f.over {
			for i := range bonus {
				if bonus[i] > 0 {
					bonus[i]--
					result += s
				}
			}
		}
	}
	return result
}
