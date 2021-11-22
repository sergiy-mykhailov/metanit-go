package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("---- Форматируемый ввод ----")

	fmt.Println("---- v1:")
	var name1 string
	var age1 int
	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name1)
	fmt.Print("Введите возраст: ")
	fmt.Fscan(os.Stdin, &age1)
	fmt.Println("result:", name1, age1)

	fmt.Println("---- v2:")
	var name2 string
	var age2 int
	fmt.Print("Введите имя: ")
	fmt.Scan(&name2)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age2)
	fmt.Println("result:", name2, age2)

	fmt.Println("---- v3:")
	var name3 string
	var age3 int
	fmt.Print("Введите имя, возраст: ")
	fmt.Scanf("%s %d", &name3, &age3)
	fmt.Println("result:", name3, age3)
}
