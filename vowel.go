package main
import (
    "fmt"
    "strings"
    )

func main() {
  count :=0
  str := "Hello World ue"
  vowel := "aeiou"
  for _, char := range str{
      if strings.ContainsRune(vowel, char){
          count++
      }
  }
  fmt.Println(count)
  
}