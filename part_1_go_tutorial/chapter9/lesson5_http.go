package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("---- Отправка запросов по HTTP ----")

	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
