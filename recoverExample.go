package main

import "fmt"

func recoverFromPanic(){
    if r:=recover(); r!=nil{
        fmt.Println("Recoverd from panic:",r)
    }
}

func callingPanic(val int){
    defer recoverFromPanic()
    
    fmt.Println("staring calling panic function")
    if val < 0{
        panic("Value cannot be negative")
    }
    fmt.Println("Value is:",val)
}

func main(){
    fmt.Println("Staring in main...")
    callingPanic(-5)
    fmt.Println("End of main")
}