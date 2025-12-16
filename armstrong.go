package main

import (
	"fmt"
	"strconv"
)

func intPow(base, power uint) uint {
	if power == 0 {
		return 1
	}
	var result uint = 1
	for i := uint(0); i < power; i++ {
		result *= base
	}
	return result
}

func main() {
	var n uint

	fmt.Println("Enter a number")
	fmt.Scan(&n)

	s := strconv.FormatUint(uint64(n), 10)
	length := uint(len(s))
	t := n
	var sum uint = 0

	for n > 0 {
		sum = sum + intPow(n%10, length)
		n = n / 10
	}

	if sum == t {
		fmt.Printf("%d is Armstrong number", t)
	} else {
		fmt.Printf("%d is not an armstrong number", t)
	}
}
