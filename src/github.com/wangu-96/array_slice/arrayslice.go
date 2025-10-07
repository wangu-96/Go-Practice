package main

import (
	"fmt"
)

func main() {

	var array1 [2]string

	array1[0] = "Hello"
	array1[1] = "World"

	fmt.Println(array1)

	fmt.Println("The length of the array is:: ", len(array1))

	fmt.Println(array1[0])
	fmt.Println(array1[1])

	slice1 := []string{"Hello", "World", "from", "a", "slice"}

	fmt.Println(slice1)

	fmt.Println("The length of the slice is:: ", len(slice1))
	fmt.Println("between 0 and 3 is:: ", slice1[0:3])

}
