/*
	Write a program that prints the numbers from 1 to 100.
	But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”.
	For numbers which are multiples of both three and five print “FizzBuzz”
*/
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	cFizz     = "Fizz"
	cBuzz     = "Buzz"
	cFizzBuzz = "FizzBuzz"
)

// isFizz decides if a number is an fizz
func isFizz(i int) bool {
	return i%3 == 0
}

// isFizz decides if a number is an buzz
func isBuzz(i int) bool {
	return i%5 == 0
}

// fizzBuzz decides how to interpret the given number (like fizz or buzz).
// if it is not fizz and buzz it returns the string representation of the number
func fizzBuzz(i int) string {
	if isFizz(i) && isBuzz(i) {
		return cFizzBuzz
	}
	if isFizz(i) {
		return cFizz
	}
	if isBuzz(i) {
		return cBuzz
	}
	return strconv.Itoa(i)
}

// fizzBuzzRepeat executes the task logic and feeds the results to the Writer
//
// to implement asynchronous logic, you can use channels, and to make the logic even more flexible,
// you can use callback, but this is not the GO-way
func fizzBuzzRepeat(hiRange int, w io.Writer) {
	for i := 1; i <= hiRange; i++ {
		fmt.Fprintln(w, fizzBuzz(i))
	}
}

func main() {
	const rangeI = 100
	fizzBuzzRepeat(rangeI, os.Stdout)
}
