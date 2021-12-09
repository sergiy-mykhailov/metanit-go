package main

import (
	"fmt"
)

func main() {
	fmt.Println("---- channels ----")

	var intCh = make(chan int, 3)
	go someFunction2(3, intCh)
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)

	fmt.Println("------")
	var intCh2 = make(chan int, 3)
	go someFunction2(3, intCh2)
	fmt.Println(<-intCh2)
	fmt.Println("cap", cap(intCh2))
	fmt.Println("len", len(intCh2))

	fmt.Println("------")
	var intCh3 = make(chan int, 3)
	go someFunction2(2, intCh3)
	fmt.Println(<-intCh3)
	fmt.Println(<-intCh3)
	//fmt.Println(<-intCh2) // err

	fmt.Println("------")
	var intCh4 = make(chan int, 2)
	go someFunction2(3, intCh4)
	fmt.Println(<-intCh4)
	fmt.Println(<-intCh4)
	fmt.Println(<-intCh4)

	fmt.Println("The End")
}

func someFunction2(n int, ch chan int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}
	//ch <- 1
	//ch <- 2
	//ch <- 3
}
