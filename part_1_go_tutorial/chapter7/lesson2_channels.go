package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---- channels ----")

	var intCh = make(chan int)

	go func() {
		fmt.Println("Go routine starts")
		time.Sleep(2 * time.Second)
		intCh <- 5 // блокировка, пока данные не будут получены функцией main
		time.Sleep(1 * time.Second)
		fmt.Println("routine is blocked ?")
	}()

	//intCh <- 5
	var value1 int = <-intCh
	fmt.Println("get - data from channel", value1)
	time.Sleep(3 * time.Second)
}
