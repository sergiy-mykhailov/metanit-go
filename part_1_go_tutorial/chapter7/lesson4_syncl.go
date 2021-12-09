package main

import (
	"fmt"
)

func main() {

	results := make(map[int]int)
	structCh := make(chan struct{})
	fmt.Println("results", results)

	go factorial74(5, structCh, results)

	<-structCh // ожидаем закрытия канала structCh

	fmt.Println("results", results)
}

func factorial74(n int, ch chan struct{}, results map[int]int) {
	defer close(ch) // закрываем канал после завершения горутины
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		results[i] = result
	}
}
