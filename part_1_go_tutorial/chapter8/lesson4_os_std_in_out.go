package main

import (
	"fmt"
	"io"
	"os"
)

type phoneReader2 string

func (p phoneReader2) Read(bs []byte) (int, error) {
	count := 0
	for i := 0; i < len(p); i++ {
		if p[i] >= '0' && p[i] <= '9' {
			bs[count] = p[i]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	fmt.Println("---- Стандартные потоки ввода-вывода и io.Copy ----")

	file, err := os.Open("chapter8/lesson3.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
	fmt.Println()

	phone1 := phoneReader2("+1(234)567 90-10")
	io.Copy(os.Stdout, phone1)
	fmt.Println()
}
