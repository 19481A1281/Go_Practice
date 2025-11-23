package main
import "fmt"

func main(){

var number int

fmt.Print("Enter number: ")
fmt.Scan(&number)

var factorial=1;
for i:=2; i<= number; i++ {

factorial = factorial * i
} 

fmt.Println(factorial);
}