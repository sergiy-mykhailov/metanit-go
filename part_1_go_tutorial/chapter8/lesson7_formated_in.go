package main

import (
	"fmt"
	"io"
	"os"
)

type person87 struct {
	name   string
	age    int32
	weight float64
}

func main() {
	fmt.Println("---- Форматируемый ввод ----")

	filename := "chapter8/lesson7.txt"
	writeData(filename)
	readData(filename)
}

func writeData(filename string) {
	var people = []person87{
		{"Tom", 24, 68.5},
		{"Bob", 25, 64.2},
		{"Sam", 27, 73.6},
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, p := range people {
		fmt.Fprintf(file, "%s %d %.2f\n", p.name, p.age, p.weight)
	}
}

func readData(filename string) {
	peoples := []person87{}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for {
		var pers = person87{}
		_, err = fmt.Fscanf(file, "%s %d %f\n", &pers.name, &pers.age, &pers.weight)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		peoples = append(peoples, pers)
		//fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
	}

	for _, value := range peoples {
		fmt.Printf("%-8s %-8d %-8.2f\n", value.name, value.age, value.weight)
	}
}
