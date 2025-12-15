package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	temp := n
	reversed := 0

	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}

	if reversed == temp {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}
}
