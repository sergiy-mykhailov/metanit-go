package main

import "fmt"

func main() {
	var hello string = "hello go"
	fmt.Println(hello)

	var (
		name string = "Tom"
		age  int    = 27
	)
	fmt.Println(name)
	fmt.Println(age)

	city := "odessa"
	fmt.Println(city)
}
