package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products/{id:[0-9]+}", productsHandler)
	router.HandleFunc("/articles/{id:[0-9]+}", productsHandler)
	router.HandleFunc("/products/{category}/{id:[0-9]+}", productsHandler)
	router.HandleFunc("/", indexHandler)

	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

func productsHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	category := vars["category"]
	response := fmt.Sprintf("Product category: %s ID: %s", category, id)
	fmt.Fprint(res, response)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Index Page")
}
