package main

import (
	"fmt"
	"sync"
)

var counter76 int = 0

func main() {
	fmt.Println("---- mutex ----")

	ch := make(chan bool) // канал
	var mutex sync.Mutex  // определяем мьютекс

	for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}
	// ожидаем завершения всех горутин
	for i := 1; i < 5; i++ {
		<-ch
	}
	fmt.Println("The End")

}

func work(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock()
	counter76 = 0
	for k := 1; k <= 5; k++ {
		counter76++
		fmt.Println("Goroutine", number, "-", counter76)
	}
	mutex.Unlock()
	ch <- true
}
