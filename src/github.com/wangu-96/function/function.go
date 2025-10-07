package main

import (
	"fmt"
)

func addTwoNumbers(a int, b int) int {
	return a + b
}

func myName(name string) string {
	return name
}

func main() {

	fmt.Println(addTwoNumbers(3, 5))
	fmt.Println(myName("Wangu"))

}
