package main

import (
	"fmt"
	"os"
)

type person85 struct {
	name   string
	age    int32
	weight float64
}

func main() {
	fmt.Println("---- Форматированный вывод ----")

	file, err := os.Create("chapter8/lesson5.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprint(file, "Сегодня ")
	fmt.Fprintln(file, "хорошая погода ")
	fmt.Println()

	tom := person85{
		name:   "Tom",
		age:    24,
		weight: 68.5,
	}
	file2, err2 := os.Create("chapter8/lesson5_person.dat")
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}
	defer file2.Close()
	fmt.Fprintf(
		file2,
		"%-10s %-10d %-10.3f\n",
		tom.name, tom.age, tom.weight)
}
