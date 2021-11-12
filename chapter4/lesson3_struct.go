package main

import (
	"fmt"
)

// Структуры

type contact struct {
	email string
	phone string
}

type person3 struct {
	age  uint
	name string
	contact
}

type node3 struct {
	value int
	next  *node3
}

//func printValue(n *node3)  {
//	fmt.Println(n.value)
//	if n != nil {
//		printValue(n.next)
//	}
//}

func main() {
	fmt.Println("Структуры")

	var valera = person3{
		name: "Valera",
		age:  24,
		contact: contact{
			email: "tom@gmail.com",
			phone: "+1234567899",
		},
	}
	fmt.Println(valera)
	fmt.Println(valera.email, valera.contact.email)

	first := node3{value: 4}
	second := node3{value: 5}
	third := node3{value: 6}

	first.next = &second
	second.next = &third

	var current *node3 = &first
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
