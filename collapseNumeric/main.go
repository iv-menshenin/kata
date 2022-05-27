package collapseNumeric

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

/*
	You are given a sequence of integers. You need to represent it in a string form using comma-separated intervals.
	Example:
		given:    []int{12, 2, 11, 3, 44, 4, 5, 6, 13}
		returned: "2-6,11-13,44"
*/

func collapseNumericSequenceOOP(n []int) string {
	var s = scroller{}
	return s.scrollUp(n)
}

func collapseNumericSequenceIter(n []int) (result string) {
	switch len(n) {
	case 0:
		return
	case 1:
		return strconv.Itoa(n[0])
	}
	sort.Ints(n)
	var start = n[0]
	for pos, curr := range n {
		var sequence bool
		var lastItem = pos == len(n)-1
		if !lastItem {
			sequence = curr+1 == n[pos+1]
		}
		if !sequence {
			if len(result) > 0 {
				result += ","
			}
			if start != curr {
				result += fmt.Sprintf("%d-%d", start, curr)
			} else {
				result += fmt.Sprintf("%d", start)
			}
			if !lastItem {
				start = n[pos+1]
			}
		}
	}
	return result
}

func collapseNumericSequenceOptimistic(n []int) string {
	sort.Ints(n)
	var start, end int
	var result = bytes.NewBufferString("")
	for {
		if start > len(n)-1 {
			break
		}
		for end = start + 1; end < len(n); end++ {
			if n[end]-n[start] != end-start {
				break
			}
		}
		if result.Len() > 0 {
			result.WriteRune(',')
		}
		if a, b := n[start], n[end-1]; a != b {
			result.WriteString(strconv.Itoa(a) + "-" + strconv.Itoa(b))
		} else {
			result.WriteString(strconv.Itoa(b))
		}
		start = end
	}
	return result.String()
}
