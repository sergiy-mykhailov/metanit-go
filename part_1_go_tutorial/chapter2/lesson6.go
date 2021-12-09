package main

import "fmt"

func main() {
	var a int = 8
	var b int = 3
	//var c bool = a == b
	fmt.Println(a == b) // false
	fmt.Println(a > b)
	fmt.Println(a != b)

	fmt.Println(!(a == b))      // true
	fmt.Println(4 > 5 && 6 > 8) // false

}
