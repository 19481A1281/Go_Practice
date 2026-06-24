package main

import "fmt"

func FindIndex[T comparable](slice []T, target T) int {
	for i, item := range slice {
		if item == target {
			return i
		}
	}
	return -1
}

func main() {
	slice := []string{"Volvo", "BMW", "BENZ"}
	fmt.Println(FindIndex(slice, "BENZ"))

	slice2 := []int{94, 81, 54, 99}
	fmt.Println(FindIndex(slice2, 81))

	slice3 := []float64{94.5, 81.2, 54.78, 99.65}
	fmt.Println(FindIndex(slice3, 99.67))
}
