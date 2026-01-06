package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var w float64
	var h float64

	fmt.Print("Enter width: ")
	fmt.Scan(&w)

	fmt.Print("Enter height: ")
	fmt.Scan(&h)

	rect := Rectangle{Width: w, Height: h}
	fmt.Println("Area of rectangle:", rect.Area())

	fmt.Printf("Height: %f, Width: %f", h, w)
}
