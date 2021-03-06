package main

import (
	"fmt"
)

func main(){
	type Days int
	const (
		Sun Days = iota
		Mon
		Tue
		Wed
		Thur
		Fri
		Sat
	)

	fmt.Println(Sun, Mon, Tue, Wed, Thur, Fri, Sat)

}