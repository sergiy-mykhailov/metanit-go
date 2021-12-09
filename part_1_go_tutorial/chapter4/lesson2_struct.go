package main

import (
	"fmt"
)

// Структуры

type person2 struct {
	age  uint
	name string
}

func main() {
	fmt.Println("Структуры")

	var ivan person2 = person2{25, "Ivan"}
	var alice person2 = person2{age: 23, name: "Alice"}
	var tom = person2{name: "Tom", age: 24}
	bob := person2{name: "Bob", age: 31}
	emptyPerson := person2{}
	fmt.Println(ivan, alice, tom, bob, emptyPerson)

	var tomPointer *person2 = &tom

	fmt.Println(tom.name) // Tom
	tom.age = 38          // изменяем значение
	fmt.Println(tom)      // Tom 38

	fmt.Println(tomPointer)
	(*tomPointer).age = 32
	fmt.Println(tom, tomPointer)

	var agePointer *uint = &alice.age
	fmt.Println(agePointer)
	*agePointer = 33
	fmt.Println(alice)
}
