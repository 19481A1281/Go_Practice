package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("divide by zero")
	}
	return x / y, nil
}

func main() {

	result, err := divide(10.20, 2.0)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result: ", result)

	result, err = divide(10.0, 0.0)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result: ", result)
}
