package main

import "fmt"

func getCoordinates() (int, int, int) {
	return 70, 20, 56
}

func main() {
	x, y, _ := getCoordinates()
	fmt.Println("X:", x, "Y:", y)
}
