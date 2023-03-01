package collapseNumeric

import (
	"fmt"
	"sort"
	"strings"
)

type scroller struct {
	// nums is a source slice
	nums []int
	// pos is the current item position in the slice
	pos int
	// num represented the item value on pos position
	num int
	// start represented the first value in a continuous sequence
	start int
	// last represented the last value in a continuous sequence
	last int
	// results collects continuous sequences in string format
	results []string
}

func (s *scroller) scrollUp(n []int) string {
	sort.Ints(n)
	s.nums = n
	if !s.init() {
		return ""
	}
	for s.next() {
		if s.isBreakSequence() {
			s.flush()
		}
	}
	s.flush()
	return s.getResults()
}

func (s *scroller) init() bool {
	if len(s.nums) == 0 {
		return false
	}
	s.pos = 0
	s.num = s.nums[s.pos]
	s.start = s.num
	return true
}

func (s *scroller) next() bool {
	s.last = s.num
	s.pos++
	if s.pos < len(s.nums) {
		s.num = s.nums[s.pos]
		return true
	}
	return false
}

func (s *scroller) isBreakSequence() bool {
	return s.last+1 != s.num
}

func (s *scroller) flush() {
	if s.start == s.last {
		s.results = append(s.results, fmt.Sprintf("%d", s.start))
	} else {
		s.results = append(s.results, fmt.Sprintf("%d-%d", s.start, s.last))
	}
	s.start = s.num
}

func (s *scroller) getResults() string {
	return strings.Join(s.results, ",")
}
