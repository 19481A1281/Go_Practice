package main
import "fmt"

func main() {
  fmt.Println("Try programiz.pro")
  var n int
  fmt.Print("Enter a number:")
  fmt.Scanln(&n)
  
  for i:=1;i<=n;i++{
      for j:=i+1;j<=n;j++{
          if((i+j)%2==0){
              fmt.Println(i,j,i+j)
          }
      }
  }
}