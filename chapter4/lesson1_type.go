package main

import (
	"fmt"
)

// Объявление типов

type mile int
type library []string

func main() {
	fmt.Println("Объявление типов")

	var distance mile = 5
	fmt.Println("distance:", distance)
	distance++
	fmt.Println("distance:", distance)

	var myLibrary library = library{"Book1", "Book2", "Book3"}
	printLib1(myLibrary)
}

func printLib1(lib library) {
	for i, value := range lib {
		fmt.Print(i, "-", value, " ")
	}
}
