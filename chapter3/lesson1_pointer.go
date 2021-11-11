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
	x = 5
	fmt.Println("Address2:", p)
	fmt.Println("Value2:", *p)

	*p = 25
	fmt.Println("Value3:", *p, x)

	var pf *float64
	fmt.Println("Value:", pf)
	//fmt.Println("Value:", *pf)  // ! ошибка, указатель не указывает на какой-либо объект

	g := new(int)
	fmt.Println("Value g:", *g) // Value: 0 - значение по умолчанию
	*g = 8                      // изменяем значение
	fmt.Println("Value g:", *g) // Value: 8
}
