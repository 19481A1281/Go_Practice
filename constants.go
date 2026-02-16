package main

import "fmt"

func main() {
	const (
		A = iota
		B
		C
		D
	)
	fmt.Println("A:", A)
	fmt.Println("B:", B)
	fmt.Println("C:", C)
	fmt.Println("D:", D)

	const (
		E = 1 * (10 * iota)
		_
		G
		H
	)
	fmt.Println("E:", E)
	//fmt.Println("F:", F)
	fmt.Println("G:", G)
	fmt.Println("H:", H)

	const (
		SUNDAY = iota
		MONDAY
		TUESDAY
		WEDNESDAY
		THURSDAY
		FRIDAY
		SATURDAY
	)

	fmt.Println("Sunday:", SUNDAY)
	fmt.Println("Monday:", MONDAY)
	fmt.Println("Tuesday:", TUESDAY)
	fmt.Println("Wednesday:", WEDNESDAY)
	fmt.Println("Thursday:", THURSDAY)
	fmt.Println("Friday:", FRIDAY)
	fmt.Println("Saturday:", SATURDAY)

}
