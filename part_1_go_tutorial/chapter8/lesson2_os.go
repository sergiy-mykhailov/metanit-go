package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("---- Создание и открытие файлов ----")

	file, err := os.Create("chapter8/lesson2.txt")

	fmt.Println(file, err)

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	fmt.Println(file.Name())
	fmt.Println("---- ----")
	fmt.Println("current pid", os.Getpid())
	fmt.Print("Hostname ")
	fmt.Println(os.Hostname())
	fmt.Println("Temp dir", os.TempDir())

}
