package main

import (
	"fmt"
)

func intPow(base int, pow int) int {
	res := 1
	for i := 0; i < pow; i++ {
		res *= base
	}

	return res
}

func main() {
	var n int

	fmt.Print("Enter a number:")
	fmt.Scan(&n)

	var sum int = 0
	t := n

	for n > 0 {
		sum = sum + (n % 10)
		n /= 10
	}

	if intPow(sum, 3) == t {
		fmt.Println("Dudency number")
	} else {
		fmt.Println("Not a Dudency number")
	}
}
