package main

import "fmt"

// Срезы

func main() {
	fmt.Println("Срезы")

	var users1 [3]string = [3]string{"Tom", "Alice", "Kate"} // array
	var users2 []string = []string{"Tom", "Alice", "Kate"}   // slice

	fmt.Println(users1)

	users2[2] = "Katherine"
	fmt.Println(users2)

	var users3 = make([]string, 3)
	users3[1] = "Alice"
	fmt.Println(users3)

	users3 = append(users3, "Valery")
	fmt.Println(users3)

	fmt.Println("Оператор среза s[i:j]")
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	fmt.Println(initialUsers[:])
	fmt.Println(initialUsers[2:])
	fmt.Println(initialUsers[2:4])
	fmt.Println(initialUsers[:4])

	fmt.Println("Удаление элемента")
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	//удаляем 4-й элемент
	var n = 3
	users = append(users[:n], users[n+1:]...)
	fmt.Println(users)
}
