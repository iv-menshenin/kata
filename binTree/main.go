package binTree

type (
	intTree struct {
		heap []int
	}
)

func New(init []int) *intTree {
	var tree = intTree{heap: init}
	if tree.Len() > 0 {
		tree.makeBalance()
	}
	return &tree
}

func (t *intTree) makeBalance() {
	var n = 0
	for idxOfRightChild(n) < t.Len() {
		n = idxOfRightChild(n)
	}
	for ; n >= 0; n-- {
		t.bDown(n)
	}
}

func (t *intTree) Len() int {
	return len(t.heap)
}

func (t *intTree) Put(nodes ...int) {
	for _, node := range nodes {
		heapLen := t.Len()
		t.heap = append(t.heap, node)
		t.bUp(heapLen)
	}
}

func (t *intTree) PopMax() (int, bool) {
	heapLen := t.Len()
	if heapLen == 0 {
		return 0, false
	}

	rootNodeVal := t.heap[0]
	t.heap[0] = t.heap[heapLen-1]
	t.heap = t.heap[:heapLen-1]
	t.bDown(0)
	return rootNodeVal, true
}

func (t *intTree) bUp(currentIdx int) {
	for currentIdx > 0 {
		parentIdx := idxOfParent(currentIdx)
		if t.heap[parentIdx] < t.heap[currentIdx] {
			t.swap(currentIdx, parentIdx)
			currentIdx = parentIdx
			continue
		}
		return
	}
}

func (t *intTree) bDown(currentIdx int) {
	for heapLen := t.Len(); currentIdx < heapLen; {
		var (
			leftChildIdx              = idxOfLeftChild(currentIdx)
			rightChildIdx             = idxOfRightChild(currentIdx)
			isLeftChildExists         = leftChildIdx < heapLen
			isRightChildExists        = rightChildIdx < heapLen
			isLeftGreaterThanCurrent  = isLeftChildExists && t.heap[currentIdx] < t.heap[leftChildIdx]
			isRightGreaterThanCurrent = isRightChildExists && t.heap[currentIdx] < t.heap[rightChildIdx]
			isRightGreaterThanLeft    = isRightChildExists && t.heap[leftChildIdx] < t.heap[rightChildIdx]
		)
		switch {

		case isLeftGreaterThanCurrent && !isRightGreaterThanLeft:
			t.swap(currentIdx, leftChildIdx)
			currentIdx = leftChildIdx
			continue

		case isRightGreaterThanCurrent:
			t.swap(currentIdx, rightChildIdx)
			currentIdx = rightChildIdx
			continue
		}
		return
	}
}

func (t *intTree) swap(a, b int) {
	t.heap[a], t.heap[b] = t.heap[b], t.heap[a]
}

func idxOfLeftChild(idx int) int {
	return idx*2 + 1
}

func idxOfRightChild(idx int) int {
	return idx*2 + 2
}

func idxOfParent(idx int) int {
	return (idx - 1) / 2
}
