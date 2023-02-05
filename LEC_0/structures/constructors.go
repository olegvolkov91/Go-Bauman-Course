package main

import "fmt"

type Transmission string

const (
	Auto     Transmission = "auto"
	Mechanic Transmission = "mechanic"
)

type Vehicle struct {
	wheels                     int
	color, engine, model, name string
	transmission               Transmission
}

func NewVehicle(wheels int, color string, engine string, model string, name string, transmission Transmission) *Vehicle {
	return &Vehicle{wheels: wheels, color: color, engine: engine, model: model, name: name, transmission: transmission}
}

func (v *Vehicle) GetFullInfo() string {
	return fmt.Sprintf("\nColor: %s\nEngine: %s\nName: %s\nModel: %s\nTransmission: %s\n", v.color, v.engine, v.name, v.model, v.transmission)
}

func main() {
	car := NewVehicle(4, "red", "1.2v3", "c3", "citroen", Mechanic)
	fmt.Println(car.GetFullInfo())
}
