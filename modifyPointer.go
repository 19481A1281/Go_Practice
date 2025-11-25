package main

import "fmt"

func modifyPointer(x *int) {
	*x = 20
}

func modifyValue(x int) {
	x = 20
}

func main() {
	x := 10
	fmt.Println(x)

	modifyValue(x)
	fmt.Println(x)

	modifyPointer(&x)
	fmt.Println(x)
}
