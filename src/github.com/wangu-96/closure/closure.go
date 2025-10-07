package main

import (
	"fmt"
)

func counter() func() int {

	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {

	myCounter := counter()

	for i := 0; i < 10; i++ {

		fmt.Println(myCounter())
	}

}
