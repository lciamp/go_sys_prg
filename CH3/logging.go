package main

import (
	"log"
)

func main() {
	x := 1
	log.Printf("log.Print() function: %d", x)

	x = x + 1
	log.Printf("log.Print() function: %d", x)

	x = x +1
	// will terminate the program
	log.Panicf("log.Panicf() function: %d", x)

	x = x + 1
	log.Printf("log.Print() function: %d", x)
}
