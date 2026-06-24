package main

import "fmt"

func Abs[T int | float64](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fmt.Println(Abs(-5))
	fmt.Println(Abs(4.7))
	fmt.Println(Abs(-7.8))
}
