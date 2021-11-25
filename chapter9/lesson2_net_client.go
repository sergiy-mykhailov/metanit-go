package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("---- Client. Обработка подключений  ----")

	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	io.Copy(os.Stdout, conn)
	fmt.Println("\ndone")
}
