package main

import "fmt"

func isAutomorphic(n int) bool {
	sq := n * n

	temp := n

	for temp > 0 {
		if temp%10 != sq%10 {
			return false
		}

		temp /= 10
		sq /= 10
	}

	return true
}

func main() {

	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	if isAutomorphic(n) {
		fmt.Printf("%d is Automorphic number", n)
	} else {
		fmt.Printf("%d is not an Automorhic number", n)
	}
}
