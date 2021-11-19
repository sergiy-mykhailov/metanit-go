package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("---- Чтение и запись файлов ----")

	text := "Hello file!"
	file, err := os.Create("chapter8/lesson3.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)
	fmt.Println("Done.")

	data := []byte("Hello Bold!")
	file2, err2 := os.Create("chapter8/lesson3.bin")
	if err2 != nil {
		fmt.Println("Unable to create file:", err2)
		os.Exit(1)
	}
	defer file2.Close()
	file2.Write(data)
	fmt.Println("Done2.")

	file3, err3 := os.Open("chapter8/lesson3.txt")
	if err3 != nil {
		fmt.Println("Unable to read file:", err3)
		os.Exit(1)
	}
	defer file3.Close()
	data3 := make([]byte, 64)
	for {
		n, err4 := file3.Read(data3)
		if err4 == io.EOF {
			break
		}
		fmt.Println(string(data3[:n]))
	}
	fmt.Println("Done3.")
}
