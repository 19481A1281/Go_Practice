package main

import "fmt"

func main() {
	v := new(int)

	fmt.Println(v)  //address
	fmt.Println(*v) //zero value of int
	fmt.Println(&v) //address of the pointer variable v itself
}
