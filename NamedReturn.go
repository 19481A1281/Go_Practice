package main

import "fmt"

func addAndSubtract(x, y int) (sum, difference int) {
	sum = x + y
	difference = x - y
	return
}

func main() {
	sum, difference := addAndSubtract(5, 2)
	fmt.Println("Sum:", sum, "Difference:", difference)
}
