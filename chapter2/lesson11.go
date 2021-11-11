package main

import "fmt"

// Функции и их параметры
func main() {
	hello()

	add(2, 5)

	var a = 8
	fmt.Println("a before: ", a)
	increment(a)
	fmt.Println("a after: ", a)

	add_multiparams(1, 2)
	add_multiparams(1, 2, 3, 4, 5)
	var numbers = []int{10, 20, 30}
	add_multiparams(numbers...)
}

func hello() {
	fmt.Println("hello !!!")
}

func add(a int, b int) {
	var result int = a + b
	fmt.Println("add: ", result)
}

func increment(x int) {
	fmt.Println("x before: ", x)
	x++
	fmt.Println("x after: ", x)
}

func add_multiparams(numbers ...int) {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	fmt.Println("sum:", sum)
}
