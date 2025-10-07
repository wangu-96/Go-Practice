package main

import (
	"fmt"
)

func main() {

	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("The address of a is %v\n", b)

	*b = 10

	fmt.Println(a)
	fmt.Println(*b)

}
