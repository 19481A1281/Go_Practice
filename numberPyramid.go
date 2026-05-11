package main
import "fmt"

func main() {
  fmt.Println("Try programiz.pro")
  var n int
  
  fmt.Print("Enter a value: ")
  fmt.Scanln(&n)
  
  for i:=1;i<=n;i++{
      for j:=0;j<n-i;j++{
          fmt.Print(" ")
      }
      for j:=1;j<=i;j++{
          fmt.Print(j," ")
      }
      fmt.Println()
  }
}