package main

import (
	"fmt"
	"math"
	"metanit_go/chapter5/computation"
)

// package

func main() {
	fmt.Println("package")

	fmt.Println("Sqrt", math.Sqrt(16))
	fmt.Println("factorial", factorial(5))
	fmt.Println("computation.factorial", computation.Factorial(6))
}
