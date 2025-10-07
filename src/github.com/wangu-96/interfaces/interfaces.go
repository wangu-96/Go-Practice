package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	Width, Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func getArea(s Shape) float64 {
	return s.Area()
}

func main() {

	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 7}

	fmt.Printf("Area of Rectangle: %.2f\n", getArea(r))
	fmt.Printf("Area of Circle: %.2f\n", getArea(c))

}
