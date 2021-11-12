package main

import "fmt"

// methods

type library4 []string
type person4 struct {
	name string
	age  int
}

func (l library4) print4() {
	for _, val := range l {
		fmt.Println(val)
	}
}

func (person person4) print41(text string) {
	fmt.Println(person.name, person.age, text)
}

func main() {
	fmt.Println("methods")

	var lib library4 = library4{"Book1", "Book2", "Book3"}
	lib.print4()

	var tom person4 = person4{"Tom", 47}
	tom.print41("hello!!")
}
