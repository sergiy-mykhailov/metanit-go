package main

import "fmt"

// defer и panic

func main() {
	fmt.Println("defer и panic")

	defer finish2()
	defer finish1()
	fmt.Println("print 1")
	fmt.Println("print 2")

	fmt.Println(divide15(3, 5))
	fmt.Println(divide15(3, 0))
	fmt.Println("finish...")
}

func finish1() {
	fmt.Println("print 3")
}

func finish2() {
	fmt.Println("print 4")
}

func divide15(a, b float32) float32 {
	if b == 0 {
		panic("Division by zero! :(")
	}

	return a / b
}
