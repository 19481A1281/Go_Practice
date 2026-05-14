package main

import "fmt"

type Car struct{
	Engine
	SteeringWheel
	Transmission
}

func (c Car) Convertable(){
	fmt.Println("Converting top")
}