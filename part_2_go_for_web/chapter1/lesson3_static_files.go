package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("part_2_go_for_web/chapter1/lesson3_static"))
	http.Handle("/", fs)

	http.HandleFunc("/contact", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Contact Page")
	})

	http.HandleFunc("/products", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "part_2_go_for_web/chapter1/lesson3_static/products.html")
	})

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		fmt.Println("Something went wrong :(", err)
	}
}
