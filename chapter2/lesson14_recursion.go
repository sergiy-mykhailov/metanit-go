package main

import "fmt"

// Рекурсивные функции

func main() {
	fmt.Println("Рекурсивные функции")

	fmt.Println("factorial:")
	fmt.Println(factorial14(0))
	fmt.Println(factorial14(3))
	fmt.Println(factorial14(5))
	fmt.Println(factorial14(10))
	fmt.Println(factorial14(30))
	fmt.Println(factorial14(65))
	fmt.Println(factorial14(300))

	fmt.Println("fibbonachi:")

	fmt.Println(fibbonachi14(3))
	fmt.Println(fibbonachi14(5))
	fmt.Println(fibbonachi14(8))
	fmt.Println(fibbonachi14(12))
	fmt.Println(fibbonachi14(30))
	fmt.Println(fibbonachi14(40))
}

func factorial14(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial14(n-1)
}

func fibbonachi14(n uint64) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fibbonachi14(n-2) + fibbonachi14(n-1)
}
