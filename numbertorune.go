package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	runeSlice := []rune(n)
	for _, r := range runeSlice {
		fmt.Println(r)
	}
}
