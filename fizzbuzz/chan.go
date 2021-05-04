package main

import (
	"fmt"
	"io"
)

func fizzBuzzChan(i <-chan int) <-chan string {
	var ch = make(chan string)
	go func() {
		for n := range i {
			ch <- fizzBuzz(n)
		}
		close(ch)
	}()
	return ch
}

func rangeGen(i, up int) <-chan int {
	var ch = make(chan int)
	go func() {
		for ; i <= up; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// fizzBuzzRepeatChan executes the task logic and feeds the results to the Writer
//
// the algorithm is implemented using channels
func fizzBuzzRepeatChan(hiRange int, w io.Writer) {
	for r := range fizzBuzzChan(rangeGen(1, hiRange)) {
		fmt.Fprintln(w, r)
	}
}
