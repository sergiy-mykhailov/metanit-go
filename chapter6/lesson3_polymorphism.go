package main

import "fmt"

// Polymorphism

type Vehicle3 interface {
	move()
}

type Car3 struct{ model string }
type Aircraft3 struct{ model string }

func (c Car3) move() {
	fmt.Println(c.model, "едет")
}
func (a Aircraft3) move() {
	fmt.Println(a.model, "летит")
}

func main() {
	fmt.Println("Polymorphism")

	tesla := Car3{"Tesla"}
	volvo := Car3{"Volvo"}
	boeing := Aircraft3{"Boeing"}

	vehicles := [...]Vehicle3{tesla, volvo, boeing}
	for _, vehicle := range vehicles {
		vehicle.move()
	}
}
