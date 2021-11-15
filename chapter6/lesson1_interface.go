package main

import "fmt"

// interface

type Vehicle interface {
	move()
}

type Car struct{}
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func drive(vehicle Vehicle) {
	vehicle.move()
}

func main() {
	fmt.Println("interface")

	var tesla Vehicle = Car{}
	var boeing Vehicle = Aircraft{}
	//tesla.move()
	//boeing.move()
	drive(tesla)
	drive(boeing)
}
