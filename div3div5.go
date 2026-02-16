package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		isDiv3 := i%3 == 0
		isDiv5 := i%5 == 0

		if isDiv3 != isDiv5 {
			fmt.Println(i)
		}
	}
}
