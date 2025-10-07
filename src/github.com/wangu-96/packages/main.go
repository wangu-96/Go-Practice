package main

import (
	"fmt"
	"math"

	"github.com/wangu-96/packages/strutil"
)

func main() {
	fmt.Println("This is the main package.")
	fmt.Println("Square root of 16 is:", math.Sqrt(16))

	fmt.Println("Value of Pi is:", math.Pi)
	fmt.Println("Value of E is:", math.E)

	fmt.Println("Reversed string of 'Hello, World!' is:", strutil.Reverse("Hello, World!"))

}
