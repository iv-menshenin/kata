package main

import (
	"fmt"
)

// Count how many numbers fit between L and R that consist of the same digit in each register.
// For example
//   between 4 and 7:    4, 5, 6, 7 (4 numbers)
//   between 10 and 100: 11, 22, 33, 44, 55, 66, 77, 88, 99 (9 numbers)

func countRepeatsInRange(l, r int64) int {
	d := countRepeats(l - 1)
	a := countRepeats(r)
	return a - d
}

func countRepeats(l int64) int {
	num := fmt.Sprintf("%d", l)
	repeats := len(num) * 9
	var max = '9'
	for i, a := range num {
		switch {
		case a < max:
			if i == 0 {
				repeats -= int(max - a)
			} else {
				repeats--
			}
			max = a
			if i > 0 {
				return repeats
			}
		case a > max:
			return repeats
		}
	}
	return repeats
}
