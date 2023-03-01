package bowling

type (
	bonus     int
	gameFrame struct {
		score []int
		over  []int
	}
	calculator struct {
		bonusStack []int
	}
	gameScores struct {
		frames []gameFrame
	}
)

const (
	regular bonus = iota
	spare
	strike

	lastFrameNum = 9
)

func (g *gameScores) Roll(score int) {
	if g.isOverScore() {
		g.lastFrame().appendOver(score)
		return
	}
	if g.isNeedNewFrame() {
		g.addToNewFrame(score)
		return
	}
	g.lastFrame().appendScore(score)
}

func (g *gameScores) isOverScore() bool {
	lastFrame := len(g.frames) - 1
	return lastFrame == lastFrameNum && g.frames[lastFrame].isCompleteFrame()
}

func (g *gameScores) isNeedNewFrame() bool {
	last := g.lastFrame()
	return last == nil || last.isCompleteFrame()
}

func (g *gameScores) lastFrame() *gameFrame {
	if len(g.frames) == 0 {
		return nil
	}
	return &g.frames[len(g.frames)-1]
}

func (g *gameScores) addToNewFrame(score int) {
	g.frames = append(g.frames, gameFrame{
		score: []int{score},
	})
}

func (g *gameScores) Score() (result int) {
	var calc calculator
	for _, currFrame := range g.frames {
		for _, score := range currFrame.score {
			result += score + calc.popBonus(score)
		}
		switch currFrame.frameBonus() {
		case spare:
			calc.putBonus(1)
		case strike:
			calc.putBonus(2)
		}
		for _, score := range currFrame.over {
			result += calc.popBonus(score)
		}
	}
	return result
}

func (c *calculator) popBonus(score int) (b int) {
	for i := range c.bonusStack {
		if c.bonusStack[i] > 0 {
			c.bonusStack[i]--
			b += score
		}
	}
	return b
}

func (c *calculator) putBonus(b int) {
	c.bonusStack = append(c.bonusStack, b)
}

func (f *gameFrame) frameBonus() bonus {
	switch true {
	case f.isSpare():
		return spare
	case f.isStrike():
		return strike
	default:
		return regular
	}
}

func (f *gameFrame) isSpare() bool {
	return scoreFrame(f.score) == 10 && len(f.score) == 2
}

func (f *gameFrame) isStrike() bool {
	return scoreFrame(f.score) == 10 && len(f.score) == 1
}

func (f *gameFrame) appendScore(score int) {
	f.score = append(f.score, score)
}

func (f *gameFrame) appendOver(score int) {
	f.over = append(f.over, score)
}

func (f *gameFrame) isCompleteFrame() bool {
	return len(f.score) == 2 || scoreFrame(f.score) > 9
}

func scoreFrame(f []int) (result int) {
	for _, s := range f {
		result += s
	}
	return result
}
