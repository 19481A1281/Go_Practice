package main

import "fmt"

func main() {
	var x int = 10
	y := &x
	z := &y

	fmt.Println(x)
	fmt.Println(*y)
	fmt.Println(**z)
}
