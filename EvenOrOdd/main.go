package main

import "fmt"

type digits []int

func newInts() digits {
	ints := digits{}

	for i := range 11 {
		ints = append(ints, i)
	}

	return ints
}

func main() {
	ints := newInts()

	for _, int := range ints {
		if int%2 == 0 {
			fmt.Println(int, "even")
		} else {
			fmt.Println(int, "odd")
		}
	}
}
