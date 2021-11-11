package dijkstra

type (
	PathFinder struct {
		geo      [][]int
		explored []bool
	}
)

const (
	uncovered = 0
	infinity  = -1
)

func (p *PathFinder) GetPath(from, to int) int {
	p.explored = make([]bool, len(p.geo))
	var current = from
	var currentCost = 0
	var canContinue bool
	var uncovered []int
	for {
		p.explored[current] = true
		exploration := p.explorePath(current)
		uncovered = p.mergeExplorations(uncovered, exploration, currentCost)
		current, currentCost, canContinue = p.nextNode(uncovered)
		if !canContinue {
			break
		}
	}
	if to < len(uncovered) {
		return uncovered[to]
	}
	return infinity
}

func (p *PathFinder) mergeExplorations(old, new []int, currentCost int) []int {
	if len(old) == 0 {
		// hack: here the currentCost is always == 0
		return new
	}
	for i := 0; i < len(new); i++ {
		if new[i] == infinity {
			continue
		}
		if old[i] == infinity || old[i] > new[i]+currentCost {
			old[i] = new[i] + currentCost
		}
	}
	return old
}

func (p *PathFinder) nextNode(exploration []int) (int, int, bool) {
	const notFound = -1
	var node, weight = notFound, infinity
	for i := 0; i < len(exploration); i++ {
		if p.explored[i] || exploration[i] == infinity {
			continue
		}
		if exploration[i] < weight || node < 0 {
			weight = exploration[i]
			node = i
		}
	}
	return node, weight, node != notFound
}

func (p *PathFinder) explorePath(from int) []int {
	var exploration = make([]int, len(p.geo[from]))
	for i := 0; i < len(p.geo[from]); i++ {
		if p.geo[from][i] != uncovered {
			exploration[i] = p.geo[from][i]
			continue
		}
		exploration[i] = infinity
	}
	return exploration
}

func New(geo [][]int) *PathFinder {
	if len(geo) == 0 {
		panic("matrix is empty")
	}
	if len(geo) != len(geo[0]) {
		panic("height != width")
	}
	return &PathFinder{
		geo: geo,
	}
}
