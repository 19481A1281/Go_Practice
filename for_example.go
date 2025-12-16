package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 5

	for j < 10 {
		fmt.Println(j)
		j++
	}

	numbers := []int{10, 20, 30, 40, 50}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value %d\n", index, value)
	}
}
