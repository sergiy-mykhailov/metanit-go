package main

import (
	"fmt"
)

func main() {
	fmt.Println("---- channels ----")

	fmt.Println("Start")
	// создание канала и получение из него данных
	//fmt.Println(<-createChan1(5)) // err: блокировка
	fmt.Println(<-createChan2(5))
	fmt.Println("End")
}

func createChan1(n int) chan int {
	ch := make(chan int) // создаем канал
	ch <- n              // отправляем данные в канал
	return ch            // возвращаем канал
}

func createChan2(n int) chan int {
	ch := make(chan int) // создаем канал
	go func() {
		ch <- n // отправляем данные в канал
	}() // запускаем горутину
	return ch // возвращаем канал
}
