package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("---- Взаимодействие клиента и сервера. Client. ----")

	conn, err1 := net.Dial("tcp", "127.0.0.1:4545")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer conn.Close()

	for {
		var source string

		fmt.Print("Введите слово: ")
		_, err2 := fmt.Scanln(&source)
		if err2 != nil {
			fmt.Println("Некорректный ввод", err2)
			continue
		}

		n3, err3 := conn.Write([]byte(source))
		if n3 == 0 || err3 != nil {
			fmt.Println("write error :(", err3)
			return
		}

		fmt.Print("Перевод:")
		buff := make([]byte, 1024)
		n4, err4 := conn.Read(buff)
		if err4 != nil {
			break
		}

		fmt.Print(string(buff[0:n4]))
		fmt.Println()
	}
}
