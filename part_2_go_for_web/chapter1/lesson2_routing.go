package main

import (
	"fmt"
	"net/http"
)

type httpHandler struct {
	message string
}

func (h httpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, h.message)
}

func main() {
	// variant 1
	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "part_2_go_for_web/chapter1/lesson2.html")
	})

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Index Page")
	})

	// variant 2
	aboutHandler := httpHandler{message: "About"}
	http.Handle("/about", aboutHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
