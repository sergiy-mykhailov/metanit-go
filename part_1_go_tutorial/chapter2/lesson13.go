package main

import "fmt"

// Анонимные функции

func main() {
	fmt.Println("Анонимные функции")
	f := func(x, y int) int { return x + y }
	fmt.Println("sum: ", f(2, 5))
	fmt.Println("sum: ", f(2, 4))

	fmt.Println("Анонимная функция как аргумент функции")
	action13(10, 5, func(x, y int) int { return x + y })
	action13(10, 5, func(x int, y int) int { return x * y })

	fmt.Println("Анонимная функция как результат функции")
	fmt.Println(selectFn13(1)(2, 3))
	fmt.Println(selectFn13(2)(4, 5))
	fmt.Println(selectFn13(3)(7, 6))

	fmt.Println("Доступ к окружению")
	f2 := square13()
	fmt.Println(f2()) // 9
	fmt.Println(f2()) // 16
	fmt.Println(f2()) // 25
}

func action13(n1 int, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}

func selectFn13(n int) func(int, int) int {
	if n == 1 {
		return func(x int, y int) int { return x + y }
	} else if n == 2 {
		return func(x int, y int) int { return x - y }
	} else {
		return func(x int, y int) int { return x * y }
	}
}

func square13() func() int {
	var x int = 2
	return func() int {
		x++
		return x * x
	}
}
