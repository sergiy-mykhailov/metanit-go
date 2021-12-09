package main

import (
	"fmt"
)

func main() {
	fmt.Println("---- close channels ----")

	intCh := make(chan int, 3)
	intCh <- 10
	intCh <- 14
	close(intCh)

	for i := 0; i < cap(intCh); i++ {
		if val, opened := <-intCh; opened {
			fmt.Println(val)
		} else {
			fmt.Println("Channel closed!")
		}
	}
}
