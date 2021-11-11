package dijkstra

type (
	PathFinder struct {
		geo      [][]int
		paths    [][]int
		explored []bool
	}
)

const (
	uncovered = 0
	infinity  = -1
)

func (p *PathFinder) GetPath(from, to int) (int, []int) {
	p.explored = make([]bool, len(p.geo))
	for i := 0; i < len(p.paths); i++ {
		p.paths[i] = make([]int, 0, len(p.geo))
	}
	var current = from
	var currentCost = 0
	var canContinue bool
	var uncovered []int
	for {
		p.explored[current] = true
		exploration := p.explorePath(current)
		uncovered = p.mergeExplorations(uncovered, exploration, currentCost, current)
		current, currentCost, canContinue = p.nextNode(uncovered)
		if !canContinue {
			break
		}
	}
	if to < len(uncovered) {
		return uncovered[to], p.paths[to]
	}
	return infinity, nil
}

func (p *PathFinder) mergeExplorations(old, new []int, currentCost, current int) []int {
	if len(old) == 0 {
		for i := 0; i < len(new); i++ {
			if new[i] != infinity {
				p.paths[i] = append(p.paths[i], []int{current, i}...)
			}
			new[i] += currentCost
		}
		return new
	}
	for i := 0; i < len(new); i++ {
		if new[i] == infinity {
			continue
		}
		if old[i] == infinity || old[i] > new[i]+currentCost {
			p.paths[i] = append(append(p.paths[i][:0], p.paths[current]...), i)
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
		geo:   geo,
		paths: make([][]int, len(geo)),
	}
}
