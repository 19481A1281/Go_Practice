package main
import "fmt"

func main() {
  fmt.Println("Try programiz.pro")
  var n int
  
  fmt.Print("Enter a value: ")
  fmt.Scanln(&n)
  
  for i:=n;i>0;i--{
      for j:=n-1;j>=i;j--{
          fmt.Print(" ")
      }
      for j:=i;j>=1;j--{
          fmt.Print("* ")
      }
      fmt.Println()
  }
}