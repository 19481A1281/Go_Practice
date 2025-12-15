package main

import "fmt"

func main() {

	var year uint

	fmt.Println("Enter Year: ")
	fmt.Scan(&year)

	if (year%4 == 0) && (year%100 != 0 || year%400 == 0) {
		fmt.Printf("%d is leap year \n", year)
	} else {

		fmt.Printf("%d is not a leap year \n", year)

	}

}
