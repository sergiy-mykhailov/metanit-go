package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("---- Отправка запросов ----")

	httpRequest := "GET / HTTP/1.0\n" +
		"Host: golang.org\n\n"

	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, conn)
	fmt.Println("done")
}
