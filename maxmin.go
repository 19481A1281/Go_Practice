package main

import "fmt"

func main() {
	var arr = [...]int{8,-12, 0, -7, 87, 34,-77, 90, 45, 33}

	var max int
	min := arr[0]

	fmt.Println("Max: ",max, "Min: ",min)

	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	fmt.Println("Max: ",max, "Min: ",min)

}
