package main

import (
	"fmt"
	"net/http"
	"time"
)

type msg string

func (message msg) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprint(res, message)
	if err != nil {
		fmt.Println("Response error:", err)
	}

	fmt.Println("New request:", time.Now(), " - ", req.Method, req.Host)
}

func main() {
	msgHandler := msg("Hello from Web Server in Go")
	fmt.Println("Server is listening...")

	err := http.ListenAndServe("localhost:8181", msgHandler)
	if err != nil {
		fmt.Println("Something went wrong :(", err)
	}
}
