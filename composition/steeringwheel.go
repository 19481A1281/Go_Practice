package main
import "fmt"

type SteeringWheel struct{
}

func(sw SteeringWheel) TurnLeft(){
	fmt.Println("turn left")
}

func(sw SteeringWheel) TurnRight(){
	fmt.Println("turn right")
}