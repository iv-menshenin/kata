package scrollupNumbers

import (
	"fmt"
	"sort"
	"strconv"
)

func scrollupNumbersSerial(n []int) string {
	var s = scroller{}
	return s.scrollUp(n)
}

func scrollNumbersBoilerPrint(n []int) string {
	sort.Ints(n)
	var start int
	var curr int
	var results string
	for i, num := range n {
		last := i == len(n)-1
		if i == 0 {
			start = num
			curr = num
			if last {
				return strconv.Itoa(num)
			}
			continue
		}
		isSeria := curr+1 == num
		if isSeria {
			curr = num
		}
		if last || !isSeria {
			if results != "" {
				results += ","
			}
			if start != curr {
				results += fmt.Sprintf("%d-%d", start, curr)
			} else {
				results += fmt.Sprintf("%d", start)
			}
			start = num
			curr = num
		}
		if last && !isSeria {
			results += fmt.Sprintf(",%d", num)
		}
	}
	return results
}

func scrollNumbersSimpled(n []int) (result string) {
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
