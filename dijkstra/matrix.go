package dijkstra

type (
	matrixGraph struct {
		geo [][]int
	}
	matrixExplorer struct {
		explored        []bool
		exploredWeights []int
		exploredPaths   [][]int
	}
)

const (
	uncovered = 0
	infinity  = -1
)

func (p *matrixGraph) GetPath(from, to int) (int, []int) {
	var (
		current     = from
		canContinue bool
		explorer    = newExplorer(len(p.geo))
	)
	for {
		explorer.step(p.geo[current], current)
		if current, canContinue = explorer.nextNode(); !canContinue {
			break
		}
	}
	if to < len(explorer.exploredPaths) && len(explorer.exploredPaths[to]) > 0 {
		return explorer.exploredWeights[to], explorer.exploredPaths[to]
	}
	return infinity, nil
}

func newExplorer(l int) *matrixExplorer {
	var explorer = matrixExplorer{
		explored:        make([]bool, l),
		exploredPaths:   make([][]int, l),
		exploredWeights: make([]int, l),
	}
	for i := 0; i < len(explorer.exploredPaths); i++ {
		explorer.exploredPaths[i] = make([]int, 0, l)
	}
	return &explorer
}

func (e *matrixExplorer) analysePathTo(currentNodeIdx, targetNodeIdx, weight int) {
	currentNodePathWeight := e.exploredWeights[currentNodeIdx]
	if e.exploredWeights[targetNodeIdx] == uncovered || e.exploredWeights[targetNodeIdx] > weight+currentNodePathWeight {
		emptyBuff := e.exploredPaths[targetNodeIdx][:0]
		if len(e.exploredPaths[currentNodeIdx]) > 0 {
			e.exploredPaths[targetNodeIdx] = append(append(emptyBuff, e.exploredPaths[currentNodeIdx]...), targetNodeIdx)
		} else {
			e.exploredPaths[targetNodeIdx] = append(emptyBuff, currentNodeIdx, targetNodeIdx)
		}
		e.exploredWeights[targetNodeIdx] = weight + currentNodePathWeight
	}
}

func (e *matrixExplorer) step(node []int, currentNodeIdx int) {
	e.explored[currentNodeIdx] = true
	for i, w := range node {
		if w == uncovered {
			continue
		}
		e.analysePathTo(currentNodeIdx, i, w)
	}
}

func (e *matrixExplorer) nextNode() (int, bool) {
	var node = -1
	for i, w := range e.exploredWeights {
		if e.explored[i] || w == uncovered {
			continue
		}
		if node < 0 || w < e.exploredWeights[node] {
			node = i
		}
	}
	return node, node > -1
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
