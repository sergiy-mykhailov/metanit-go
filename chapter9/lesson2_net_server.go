package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("---- Сервер. Обработка подключений ----")

	message := "Hello, I am a server"

	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("New connection", time.Now())
		conn.Write([]byte(message))
		conn.Close()
	}
}
