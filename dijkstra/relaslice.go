package dijkstra

type (
	PathWeight struct {
		from, to int
		weight   float64
	}
	Path struct {
		path []int
		cost float64
	}
	RoadMap []PathWeight
)

const maxNodes = 1024

func (r RoadMap) explorePathsFrom(path Path) []Path {
	if len(path.path) == 0 {
		panic("path slice is empty")
	}
	current := path.path[len(path.path)-1]
	var ways = make([]Path, 0)
	for _, weight := range r {
		if weight.from != current {
			continue
		}
		ways = append(ways, Path{
			path: append(append(make([]int, 0, 16), path.path...), weight.to),
			cost: path.cost + weight.weight,
		})
	}
	return ways
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
	if mNode > maxNodes {
		panic("max nodes count exceeded")
	}
	return mNode
}

func selectNextNode(uncovered []bool, explored []Path) (next int, cost float64) {
	next = -1
	for i := range uncovered {
		if uncovered[i] {
			continue
		}
		if len(explored[i].path) == 0 {
			continue
		}
		if next < 0 || cost > explored[i].cost {
			next = i
			cost = explored[i].cost
		}
	}
	return
}

func (r RoadMap) explorePathFromTo(from, to int) Path {
	var currNode = from
	var currCost float64
	var explored = make([]Path, r.getMaxNode()+1)
	var uncovered = make([]bool, r.getMaxNode()+1)
	explored[from] = Path{path: []int{currNode}, cost: currCost}
	for {
		for _, path := range r.explorePathsFrom(explored[currNode]) {
			dest := path.path[len(path.path)-1]
			if len(explored[dest].path) == 0 || explored[dest].cost > path.cost {
				explored[dest] = path
			}
		}
		uncovered[currNode] = true
		currNode, currCost = selectNextNode(uncovered, explored)
		if currNode < 0 {
			break
		}
	}
	return explored[to]
}
