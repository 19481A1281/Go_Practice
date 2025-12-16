package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i*j == 40 {
				fmt.Println("found it", i, "*", j)
				goto end
			}
			fmt.Println(i, j)
		}
	}

end:
	fmt.Println("Done searching")
}
