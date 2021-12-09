package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("---- Буферизированный ввод-вывод ----")

	fileName := "chapter8/lesson9.txt"
	writeWithBuffer(fileName)
	readWithBuffer(fileName)
}

func writeWithBuffer(fileName string) {
	rows := []string{"hello go!", "welcome to golang!"}

	file, err := os.Create(fileName)
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer file.Close()

	for _, row := range rows {
		writer.WriteString(row)
		writer.WriteString("\n")
	}

	writer.Flush()
}

func readWithBuffer(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}

		fmt.Println(line)
	}
}
