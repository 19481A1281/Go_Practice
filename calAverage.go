package main

import "fmt"

func calculateAverage(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	numbers := []float64{10.0, 23.0, 76.34}
	fmt.Print("Average: ", calculateAverage(numbers))
}
