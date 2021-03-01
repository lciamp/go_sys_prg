package main

import (
	"fmt"
)

// unnamed return values
func unnamedMinMax(x, y int) (int, int){
	if x > y {
		min := y
		max := x
		return min, max
	} else {
		min := x
		max := y
		return min, max
	}
}

// named return values
func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return min, max
}

// same as above, don't need to name the return order, it's in the def
func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}

// anonymous functions
func main() {
	y := 4

	// DO NOT GIVE 2 ANONYMOUS FUNCTIONS THE SAME NAME:
	square := func(s int) int {
		return s * s
	}

	fmt.Println("The square of", y, "is", square(y))

	square = func(s int) int {
		return s + s
	}

	fmt.Println("The square of", y, "is", square(y))

	// defined functions
	fmt.Println(minMax(15, 6))
	fmt.Println(namedMinMax(15, 6))
	min, max := namedMinMax(12, -1)
	fmt.Println(min, max)
}