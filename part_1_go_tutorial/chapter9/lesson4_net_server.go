package main

import (
	"fmt"
	"net"
	"time"
)

var dict94 = map[string]string{
	"red":    "красный",
	"green":  "зеленый",
	"blue":   "синий",
	"yellow": "желтый",
}

func main() {
	fmt.Println("---- Установка таймаута ----")

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
			conn.Close()
			continue
		}

		go handleConnection94(conn)
	}
}

func handleConnection94(conn net.Conn) {
	defer conn.Close()

	fmt.Println("New connection", time.Now())

	for {
		input := make([]byte, 1024*4)

		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}

		source := string(input[0:n])
		target, ok := dict94[source]
		if ok == false {
			target = "undefined"
		}

		fmt.Println(source, "-", target)
		conn.Write([]byte(target))
	}
}
