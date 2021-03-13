package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {

	mySlice := make([]aStructure, 0)

	a := aStructure{"lou", 180, 90}
	mySlice = append(mySlice, a)
	a = aStructure{"bean", 70, 50}
	mySlice = append(mySlice, a)
	a = aStructure{"hudson", 50, 20}
	mySlice = append(mySlice, a)
	a = aStructure{"bill", 78, 91}
	mySlice = append(mySlice, a)

	fmt.Println("0: ", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight < mySlice[j].weight
	})
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight > mySlice[j].weight
	})
	fmt.Println(">:", mySlice)
}
