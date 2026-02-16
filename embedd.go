package main
import "fmt"

type Address struct{
    HouseNo string
    City string
    Pincode string
}

type Person struct{
    Name string
    Age uint
    Address 
    //Address Address
}

func main() {
  
  person := Person{
      Name: "chin",
      Age: 24,
      Address : Address {
          HouseNo: "3-89",
          City: "Bangalore",
          Pincode: "432454",
      },
  }
  
  fmt.Println(person)
  fmt.Println(person.Address.City)
  fmt.Println(person.Pincode)
}