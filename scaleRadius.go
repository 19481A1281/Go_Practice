package main

import "fmt"

type Circle struct {
	Radius float64
}

func (c Circle) Scale(factor float64) {
	c.Radius = c.Radius * factor
	fmt.Println("Scaled Radius (inside method):", c.Radius)
}

func main() {
	var r float64
	fmt.Print("Enter radius of a circle:")
	fmt.Scan(&r)

	circle := Circle{Radius: r}

	var factor float64
	fmt.Print("Enter value to scale:")
	fmt.Scan(&factor)
	circle.Scale(factor)

	fmt.Print("Radius after scaling:", circle.Radius)
}
