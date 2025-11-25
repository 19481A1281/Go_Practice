package main

import "fmt"

func main() {
	var p *int
	fmt.Println("Value of p:", p) //<nil>

	// fmt.Println("Value pointed to by p:", *p) // Panic: invalid memory address or nil pointer dereference
}
