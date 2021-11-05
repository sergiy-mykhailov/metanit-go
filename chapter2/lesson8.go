package main

import "fmt"

func main() {
	var numbers [5]int
	fmt.Println(numbers)

	var numbers2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers2)

	var numbers3 = [...]int{1, 2, 3} // длина массива 3
	fmt.Println(numbers3)            // [1 2 3]

	colors := [3]string{2: "blue", 0: "red", 1: "green"}
	fmt.Println(colors[2]) // blue
}
