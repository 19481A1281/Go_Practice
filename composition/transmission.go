package main

import "fmt"

type Transmission struct{
}

func (t Transmission) ShiftUp(){
	fmt.Println("Shifting up")
}

func(t Transmission) ShiftDown(){
	fmt.Println("Shifting Down")
}