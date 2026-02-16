package main

import "fmt"

func main() {
	arr := []int{7, 2, 5, 3, 1, 8}
	sum := 10
abc:
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				fmt.Printf("pair found at %d and %d (%d + %d)", i, j, arr[i], arr[j])
				break abc
			}
		}
	}
}
