package main

import "fmt"

// Возвращение результата из функции
func main() {
	fmt.Println("sum: ", add1(2, 5))
	fmt.Println("sum: ", add2(2, 5))
	var age, name = add3(4, 5, "Tom", "Ivanov")
	fmt.Println("sum: ", age, name)

	var f func(int, int) int = add2
	fmt.Println("f: ", f(2, 5))

	action(10, 25, add1)   // 35
	action(5, 6, multiply) // 30

	gg := selectFn(1)
	fmt.Println(gg(3, 4)) // 7

	gg = selectFn(3)
	fmt.Println(gg(3, 4))          // 12
	fmt.Println(selectFn(2)(3, 4)) // -1

}

func add1(x, y int) int {
	return x + y
}

func add2(x, y int) (z int) {
	z = x + y
	return
}

func add3(x, y int, firstName, lastName string) (int, string) {
	var z = x + y
	var fullName = firstName + " " + lastName

	return z, fullName
}

func multiply(x int, y int) int {
	return x * y
}

func action(n1 int, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}

func subtract(x int, y int) int { return x - y }

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return add1
	} else if n == 2 {
		return subtract
	} else {
		return multiply
	}
}
