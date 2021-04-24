package main

import "io"

type (
	s struct {
		s string
		i int
	}
	a struct {
		arr []s
	}
	a1 []s
)

// awesomeFuncWithStruct uses a struct reference to logically separate the two states:
//  - collection does not exist
//  - collection exists, but no data in it
// but why, then, is the error passed to the result?
// bad and unobvious logic, creates an unnecessary load on resources. however i saw it in production
func awesomeFuncWithStruct(i int) (*a, error) {
	if i%2 == 0 {
		return nil, io.EOF
	}
	result := &a{}
	for nn := 0; nn < 10; nn++ {
		result.arr = append(result.arr, s{s: "some foo", i: nn})
	}
	return result, nil
}

// awesomeFuncWithSlice now we understand that EOF means that there is no collection,
// however, if we know in advance the amount of fetch data, we have one more optimization opportunity.
//  var result []s                       // not good
//  var result = make([]s, 0, someCount) // good
func awesomeFuncWithSlice(i int) (a1, error) {
	if i%2 == 0 {
		return nil, io.EOF
	}
	var result []s
	for nn := 0; nn < 10; nn++ {
		result = append(result, s{s: "some foo", i: nn})
	}
	return result, nil
}

// awesomeFuncWithMake uses EOF to indicate the absence of a collection and pre-specifies the size of the array
// to which the data will be added
func awesomeFuncWithMake(i int) (a1, error) {
	if i%2 == 0 {
		return nil, io.EOF
	}
	var result = make([]s, 0, 10)
	for nn := 0; nn < 10; nn++ {
		result = append(result, s{s: "some foo", i: nn})
	}
	return result, nil
}

func main() {
	println("bench me")
}
