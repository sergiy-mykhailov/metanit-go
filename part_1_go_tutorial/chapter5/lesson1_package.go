package main

import (
	"fmt"
	"math"
	computation2 "metanit_go/part_1_go_tutorial/chapter5/computation"
)

// package

func main() {
	fmt.Println("package")

	fmt.Println("Sqrt", math.Sqrt(16))
	fmt.Println("factorial", factorial(5))
	fmt.Println("computation.factorial", computation2.Factorial(6))
}
