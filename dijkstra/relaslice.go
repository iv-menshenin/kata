package dijkstra

type (
	PathWeight struct {
		from, to int
		weight   float64
	}
	Path struct {
		path     []int
		cost     float64
		explored bool
	}
	RoadMap     []PathWeight
	Exploration []Path
)

const maxNodes = 1024

func (e Exploration) explorePathsFrom(r RoadMap, from int) {
	for _, way := range r {
		if way.from != from {
			continue
		}
		cost := e[from].cost + way.weight
		if len(e[way.to].path) == 0 || e[way.to].cost > cost {
			e[way.to].path = append(append(e[way.to].path[:0], e[from].path...), way.to)
			e[way.to].cost = cost
		}
	}
}

func (r RoadMap) getMaxNode() int {
	var mNode = -1
	for _, node := range r {
		if mNode < node.from {
			mNode = node.from
		}
		if mNode < node.to {
			mNode = node.to
		}
	}
	if !(mNode < maxNodes) {
		panic("max nodes count exceeded")
	}
	return mNode
}

func (e Exploration) selectNextNode() (next int) {
	var cost float64
	next = -1
	for i := range e {
		if e[i].explored {
			continue
		}
		if len(e[i].path) == 0 {
			continue
		}
		if next < 0 || cost > e[i].cost {
			next = i
			cost = e[i].cost
		}
	}
	return
}

func (r RoadMap) explorePathFromTo(from, to int) Path {
	var currNode = from
	var exploration = newExploration(r.getMaxNode() + 1)
	exploration[from].path = append(exploration[from].path, currNode)
	for {
		exploration.explorePathsFrom(r, currNode)
		exploration[currNode].explored = true
		currNode = exploration.selectNextNode()
		if currNode < 0 {
			break
		}
	}
	return exploration[to]
}

func newExploration(dim int) Exploration {
	var ex [maxNodes]Path
	var place = make([]int, dim*dim)
	var exploration = ex[:dim]
	for i := 0; i < dim; i++ {
		exploration[i].path = (place[i*dim : i*dim+dim])[:0]
	}
	return exploration
}
