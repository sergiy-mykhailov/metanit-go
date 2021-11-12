package main

import "fmt"

// methods

type person5 struct {
	name string
	age  int
}

func (p person5) updateAge(newAge int) {
	p.age = newAge
}

func (p *person5) updateAge2(newAge int) {
	(*p).age = newAge
}

func main() {
	fmt.Println("methods and pointers")

	var tom = person5{name: "Tom", age: 24}
	fmt.Println("before", tom.age)
	tom.updateAge(33)
	fmt.Println("after", tom.age)
	tom.updateAge2(33)
	fmt.Println("after", tom.age)
}
