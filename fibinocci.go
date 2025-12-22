package main

import "fmt"

func main() {

	var n int

	fmt.Print("Enter a number to print fibinocci series: ")
	fmt.Scanln(&n)

	n1 := 0
	n2 := 1

	fmt.Print(n1, "\n", n2, "\n")

	t := 0

	for n > 0 {
		t = n1 + n2
		n1 = n2
		n2 = t
		fmt.Println(t)
		n--
	}

}
