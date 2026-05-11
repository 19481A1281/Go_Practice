package main
import "fmt"

func main() {
  fmt.Println("Try programiz.pro")
  str1:="Hello తెలుగు"
  
  for _, value :=  range str1{
      fmt.Printf("Charcater: %c ,Value : %+v , Type : %T \n",value,value,value)
  }
  fmt.Println("Bytes")
  str2 := "Hello World"
  for i:=0; i<len(str2);i++{
      fmt.Printf("Charcater: %c ,Value : %+v , Type : %T \n",str2[i],str2[i],str2[i])
  }
}