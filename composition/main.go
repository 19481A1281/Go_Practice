package main

//import "fmt"

type Startable interface{
	Start()
}

func startEngines(vehicles ...Startable){
	for _,v := range vehicles{
		v.Start()
	}
}

func main(){
	car := Car{Engine{}, SteeringWheel{}, Transmission{}}
	truck := Truck{Engine{}, Transmission{}, SteeringWheel{}}

	car.Convertable()
	truck.FourWheelDrive()

	startEngines(car, truck)
}