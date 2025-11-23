package main

import "fmt"

func reverse(n int) int{
sum :=0
for n>0 {
sum = sum*10 + (n%10)
n=n/10
}

return sum
}

func main(){

var n int
fmt.Print("Enter number: ")
fmt.Scan(&n)
fmt.Println(reverse(n))
}