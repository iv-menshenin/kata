package dijkstra

type (
	matrixGraph struct {
		geo [][]int
	}
	matrixExplorer struct {
		explored []bool
		current  []int
		paths    [][]int
	}
)

const (
	uncovered = 0
	infinity  = -1
)

func (p *matrixGraph) GetPath(from, to int) (int, []int) {
	var explorer = matrixExplorer{
		explored: make([]bool, len(p.geo)),
		paths:    make([][]int, len(p.geo)),
	}
	for i := 0; i < len(explorer.paths); i++ {
		explorer.paths[i] = make([]int, 0, len(p.geo))
	}
	var current = from
	var currentCost = 0
	var canContinue bool
	for {
		explorer.explored[current] = true
		exploration := explorer.explorePath(p.geo[current])
		explorer.mergeExplorations(exploration, currentCost, current)
		current, currentCost, canContinue = explorer.nextNode()
		if !canContinue {
			break
		}
	}
	if to < len(explorer.current) {
		return explorer.current[to], explorer.paths[to]
	}
	return infinity, nil
}

func (e *matrixExplorer) mergeExplorations(bestPaths []int, currentCost, currentIdx int) {
	if len(e.current) == 0 {
		for i := 0; i < len(bestPaths); i++ {
			if bestPaths[i] != infinity {
				e.paths[i] = append(e.paths[i], []int{currentIdx, i}...)
			}
			bestPaths[i] += currentCost
		}
		e.current = bestPaths
		return
	}
	for i := 0; i < len(bestPaths); i++ {
		if bestPaths[i] == infinity {
			continue
		}
		if e.current[i] == infinity || e.current[i] > bestPaths[i]+currentCost {
			e.paths[i] = append(append(e.paths[i][:0], e.paths[currentIdx]...), i)
			e.current[i] = bestPaths[i] + currentCost
		}
	}
}

func (e *matrixExplorer) nextNode() (int, int, bool) {
	const notFound = -1
	var node, weight = notFound, infinity
	for i := 0; i < len(e.current); i++ {
		if e.explored[i] || e.current[i] == infinity {
			continue
		}
		if e.current[i] < weight || node < 0 {
			weight = e.current[i]
			node = i
		}
	}
	return node, weight, node != notFound
}

func (e *matrixExplorer) explorePath(node []int) []int {
	var exploration = make([]int, len(node))
	for i := 0; i < len(node); i++ {
		if node[i] != uncovered {
			exploration[i] = node[i]
			continue
		}
		exploration[i] = infinity
	}
	return exploration
}

func New(geo [][]int) *matrixGraph {
	if len(geo) == 0 {
		panic("matrix is empty")
	}
	if len(geo) != len(geo[0]) {
		panic("height != width")
	}
	return &matrixGraph{
		geo: geo,
	}
}
