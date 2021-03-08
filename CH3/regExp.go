package main

import (
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("Mihalis", "Mihalis Tsoukalos")
	fmt.Println(match) // true

	match, _ = regexp.MatchString("Tsoukalos", "Mihalis tsoukalos")
	fmt.Println(match) // false

	parse, err := regexp.Compile("[Mm]ihalis")

	if err != nil {
		fmt.Printf("Error compiling RE: %s\n", err)
	} else {
		fmt.Println(parse.MatchString("Mihalis Tsoukalos")) // true
		fmt.Println(parse.MatchString("mihalis Tsoukalos")) // true
		fmt.Println(parse.MatchString("Mihalis Tsoukalos")) // true
		fmt.Println(parse.ReplaceAllString("mihalis Mihalis", "MIHALIS")) // MIHALIS MIHALIS

	}
}
