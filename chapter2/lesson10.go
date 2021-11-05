package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	var i = 1
	for ; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	var j = 1
	for j < 10 {
		fmt.Print(j*j, " ")
		j++
	}
	fmt.Println()

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	var users = [3]string{"Tom", "Alice", "Kate"}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}

	for index, value := range users {
		fmt.Println(index, value)
	}

	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	var sum = 0
	for _, value := range numbers {
		if value < 0 {
			continue // переходим к следующей итерации
		}
		if value > 5 {
			break // если число больше 4 выходим из цикла
		}
		sum += value
	}
	fmt.Println("Sum:", sum)
}
