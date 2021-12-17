package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/user", func(res http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		age := req.URL.Query().Get("age")

		fmt.Fprintf(res, "Имя: %s \nВозраст: %s", name, age)
	})

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "part_2_go_for_web/chapter1/lesson5.html")
	})

	http.HandleFunc("/postform", func(res http.ResponseWriter, req *http.Request) {

		name := req.FormValue("username")
		age := req.FormValue("userage")

		fmt.Fprintf(res, "Имя: %s Возраст: %s", name, age)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
