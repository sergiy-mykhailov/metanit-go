package main

import (
	"fmt"
)

func main() {
	fmt.Println("---- channels ----")

	var intCh = make(chan int)

	go factorial22(5, intCh) // вызов горутины
	fmt.Println(<-intCh)     // получение данных из канала
	fmt.Println("The End")
}

func factorial22(n int, ch chan int) {

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)

	ch <- result // отправка данных в канал
}
