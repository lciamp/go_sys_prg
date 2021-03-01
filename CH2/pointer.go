package main

import (
	"fmt"
)

// no return because you are directly changing the val by using a pointer
func withPointers(x *int) {
	*x = *x * *x
}


type complex struct {
	x, y int
}

func newComplex(x, y int) *complex {
	return &complex{x, y}
}

func main(){
	x := 2
	withPointers(&x)
	fmt.Println(x)

	w := newComplex(4, -5)
	fmt.Println(*w)
	fmt.Println(w)
}