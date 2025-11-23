package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")
	name := "Durga Prasad"
	fmt.Println("Hello", name+"!")
	//message := helper.Greeting("World")
	//fmt.Println(message)
	fmt.Println("Arguments: ", os.Args[1:])
}
