package main

import (
	"fmt"
	"strings"
)

func main() {
	var s [3]string
	s[0] = "1 2 3"
	s[1] = "11 12 13 14 15 16"
	s[2] = "-1 2 -3 -4 -5 6"

	column := 3

	for i := 0; i < len(s); i++ {
		data := strings.Fields(s[i]) // Fields splits a string based on whitespace (like split in python)
		if len(data) >= column {
			fmt.Println(data[column-1])
		}
	}
}
