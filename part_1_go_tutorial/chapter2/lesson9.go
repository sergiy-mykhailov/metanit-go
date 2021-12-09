package main

import "fmt"

func main() {
	a := 6
	b := 7
	if a < b {
		fmt.Println("a is less than b")
	}

	if a < b {
		fmt.Println("a is less than b - 2")
	}

	c := 8
	switch c + 2 {
	case 9:
		fmt.Println("c = 9")
	case 8:
		fmt.Println("c = 8")
	case 7:
		fmt.Println("c = 7")
	default:
		fmt.Println("значение переменной c не определено")
	}

	d := 5
	switch d {
	case 9:
		fmt.Println("d = 9")
	case 8:
		fmt.Println("d = 8")
	case 7:
		fmt.Println("d = 7")
	case 6, 5, 4:
		fmt.Println("d = 6 или 5 или 4, но это не точно")
	default:
		fmt.Println("значение переменной a не определено")
	}
}
