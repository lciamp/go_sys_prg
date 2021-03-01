// defer keyword
package main

import (
	"fmt"
)

func a1() {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")
	}
}

// anonymous function will disappear once loop is done, uses 'i' directly
func a2() {
	for i := 0; i < 3; i++ {
		defer func() {fmt.Print(i, " ")}()
	}
}

// func needs parameter 'n', gets it from 'i'
func a3() {
	for i := 0; i < 3; i++ {
		defer func(n int) { fmt.Print(n, " ")}(i)
	}
}

func main() {
	a1()
	fmt.Println()
	a2()
	fmt.Println()
	a3()
	fmt.Println()
}

// output
// 2 1 0
// 3 3 3
// 2 1 0

// defer functions are executed in LIFO after the return of surrounding function
// a1() waits for the loop to finish then prints in reverse
// a2() function is eval'd after the loop uses the where i is at: 3
// a3() gets a different i for each time the function is called


