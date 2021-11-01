package main

import "fmt"

func main() {
	//const pi float64 = 3.1415
	//pi = 2.7182 // error

	const pi, e = 3.1415, 2.7182
	fmt.Println(pi, e)

	const (
		a = 1
		b
		c
		d = 3
		f
	)
	fmt.Println(a, b, c, d, f) // 1, 1, 1, 3, 3

	var m int = 7
	// const k = m      // error
	const s = 5
	const n = s

	fmt.Println(m, s, n)
}
