package main

import (
	"fmt"
	"time"
)

// goroutines

func main() {
	fmt.Println("goroutines")

	for i := 1; i <= 16; i++ {
		go factorial(i)
	}

	go func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("anonymous function")
	}()

	//fmt.Scanln()
	time.Sleep(2 * time.Second)
	fmt.Println("The End")
}

func factorial(n int) {
	if n < 1 {
		fmt.Println("Invalid input number")
		return
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	fmt.Println(n, "-", result)
}
