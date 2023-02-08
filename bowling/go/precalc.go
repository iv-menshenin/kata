package bowling

type (
	bowling struct {
		score       int
		frameScore  int
		frameRoll   int
		singleBonus bool
		doubleBonus bool
	}
)

func New() *bowling {
	return &bowling{}
}

func (b *bowling) Roll(score int) {
	b.intRoll(score)
	b.frameStateChange()
	b.frameEnd()
}

func (b *bowling) frameStateChange() {
	if b.frameScore == 10 {
		switch b.frameRoll {
		case 1:
			b.doubleBonus = true
		case 2:
			b.singleBonus = true
		}
	}
}

func (b *bowling) frameEnd() {
	if b.frameRoll == 2 || b.frameScore == 10 {
		b.frameRoll = 0
		b.frameScore = 0
	}
}

func (b *bowling) intRoll(score int) {
	b.frameRoll++
	b.score += score
	b.frameScore += score
	b.checkSpare(score)
	b.checkStrike(score)
}

func (b *bowling) checkSpare(score int) {
	if b.singleBonus {
		b.singleBonus = false
		b.score += score
	}
}

func (b *bowling) checkStrike(score int) {
	if b.doubleBonus {
		b.singleBonus = true
		b.doubleBonus = false
		b.score += score
	}
}

func (b *bowling) Score() int {
	return b.score
}
