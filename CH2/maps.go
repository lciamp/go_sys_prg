package main

import (
	"fmt"
)

func main(){
	// basic map/dict
	aMap := make(map[string]int)

	aMap["Mon"] = 0
	aMap["Tue"] = 1
	aMap["Wed"] = 2
	aMap["Thu"] = 3
	aMap["Fri"] = 4
	aMap["Sat"] = 5
	aMap["Sun"] = 6

	fmt.Printf("Sunday is the %dth day of the week.\n", aMap["Sun"])

	// check to see if a key exists
	_, ok := aMap["Tuesday"]
	if ok {
		fmt.Printf("The Tuesday key exists!\n")
	} else {
		fmt.Printf("The Tuesday key does not exists!\n")
	}

	//iterate over all keys
	count := 0
	for key, _ := range aMap {
		count++
		fmt.Printf("%s ", key)
	}
	fmt.Printf("\n")
	fmt.Printf("The map has %d elements\n", count)

	//just count pairs
	count = 0
	delete(aMap, "Fri")
	for _, _ = range aMap {
		count ++
	}
	fmt.Printf("There are %d elements in the map.\n", count)

	//other way of defining
	anotherMap := map[string]int{
		"one" : 1,
		"two" : 2,
		"three" : 3,
		"four" : 4,
	}
	anotherMap["five"] = 5
	count = 0
	for _,_ = range anotherMap {
		count++
	}
	fmt.Printf("anotherMap has %d elements.\n", count)
}
