package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter a number:")
	fmt.Scan(&n)

	sum := 0

	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	if sum == n {
		fmt.Println("Perfect number")
	} else {
		fmt.Println("Not a perfect number")
	}
}
