package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		n uint
		e uint
		o uint
	)
	fmt.Print("Enter a number:")
	fmt.Scan(&n)

	start := time.Now()
	for n > 0 {
		if n&1 == 0 {
			e++
		} else {
			o++
		}
		n /= 10
	}
	duration := time.Since(start)

	fmt.Println("Even:", e)
	fmt.Println("Odd:", o)
	fmt.Printf("Execution Time: %v\n", duration)

}
