package main
import "fmt"

func main() {
  var n int
  fmt.Print("Enter a number: ")
  fmt.Scanln(&n)
  
  l := 0
  for n > 0{
      if n%10 > l{
          l = n%10
          if l ==9 {
              fmt.Println("Breaking loop here")
              break
          }
      }
      n/=10
  }
  
  fmt.Println("Larger digit:",l)
}