package main

import (
	"fmt"
)

func main() {
	fmt.Println("---- Передача потоков данных ----")

	intCh := make(chan int)

	go factorial75(7, intCh)

	for {
		num, opened := <-intCh // получаем данные из потока
		if !opened {
			break // если поток закрыт, выход из цикла
		}
		fmt.Println(num)
	}

	//// alternative:
	//for num := range intCh{
	//	fmt.Println(num)
	//}
}

func factorial75(n int, ch chan int) {
	defer close(ch)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		ch <- result // посылаем по числу
	}
}
