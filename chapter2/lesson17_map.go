package main

import "fmt"

// map

func main() {
	fmt.Println("map")

	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	fmt.Println(people) // map[Tom:1 Bob:2 Sam:4 Alice:8]

	fmt.Println(people["Bob"]) // 2
	people["Bob"] = 32
	fmt.Println(people["Bob"]) // 32

	if val, ok := people["Tom"]; ok {
		fmt.Println(val, ok)
	}

	var value1, present1 = people["Sam"]
	fmt.Println(value1, present1)
	var value2, present2 = people["555"]
	fmt.Println(value2, present2)

	for key, value := range people {
		fmt.Print(key, "-", value, " ")
	}
	fmt.Println()
	fmt.Println(make(map[string]int))

	people["Kate"] = 128
	fmt.Println(people)

	delete(people, "Tom")
	fmt.Println(people)
}
