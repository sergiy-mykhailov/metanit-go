package main

import "fmt"

// Указатели

func main() {
	fmt.Println("Указатели")

	var x int = 4               // определяем переменную
	var p *int                  // определяем указатель
	p = &x                      // указатель получает адрес переменной
	fmt.Println("Address1:", p) // значение самого указателя - адрес переменной x
	fmt.Println("Value1:", *p)

	d := 5
	fmt.Println("d before:", d) // 5
	changeValue1(d)             // изменяем значение
	fmt.Println("d after:", d)  // 5 - значение не изменилось

	fmt.Println("d before:", d) // 5
	changeValue2(&d)            // изменяем значение
	fmt.Println("d after:", d)  // 25 - значение изменилось!

	p1 := createPointer2(7)
	fmt.Println("p1:", *p1, p1) // p1: 7
	p2 := createPointer2(10)
	fmt.Println("p2:", *p2, p2) // p2: 10
	p3 := createPointer2(28)
	fmt.Println("p3:", *p3, p3) // p3: 28
}

func changeValue1(x int) {
	x = x * x
}

func changeValue2(x *int) {
	*x = (*x) * (*x)
}

func createPointer2(x int) *int {
	p := new(int)
	*p = x
	return p
}
