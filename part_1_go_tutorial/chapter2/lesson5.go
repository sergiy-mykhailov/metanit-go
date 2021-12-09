package main

import "fmt"

func main() {
	var a int = 10
	var b int = 4
	var c int = a / b
	fmt.Println(c) // 2

	var m float32 = 11 / 4 // 2
	fmt.Println(m)

	var n float32 = 10 / 4.0 // 2.5
	fmt.Println(n)

	var k int = 35 % 3 // 2 (35 - 33 = 2)
	fmt.Println(k)

}
