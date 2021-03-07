/*
Write a Go program that keeps reading ints until you give the number 0 inout,
then it prints the min and max integer in the input
 */
package main

import (
	"fmt"
)

func main() {
	x := 1
	max, min, count := 0, 0, 0

	for x != 0 {
		fmt.Println("Please enter a number")
		fmt.Scanf("%d", &x)

		if count == 0{
			min = x
			max = x
		}

		if x >= max {
			max = x
		} else if x > 0 && x < min{
			min = x
		}
		count++
	}
	fmt.Println("Min:", min, "Max:", max)
}