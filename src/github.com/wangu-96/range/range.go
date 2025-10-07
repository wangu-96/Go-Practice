package main

import (
	"fmt"
)

func main() {

	num := []int{2, 4, 6, 8}

	for i, num := range num {

		fmt.Printf("The index is %d and the value is %d\n", i, num)

	}

	email := map[string]string{"name": "Wangu", "email": "wangungalati@gmail.com"}

	for k, v := range email {

		fmt.Printf("The key is %s and the value is %s\n", k, v)

	}
}
