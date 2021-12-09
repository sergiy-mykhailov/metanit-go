package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type product struct {
	id      int
	model   string
	company string
	price   int
}

func main() {
	fmt.Println("---- postgres ----")

	connStr := "user=postgres password=postgres dbname=testgo sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createTable(db)

	productId := insertProduct(db)
	getProducts(db)
	updateProduct(db, productId)
	deleteProduct(db, productId)
}

func createTable(db *sql.DB) {
	fmt.Println("Creating table...")

	result, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS Products (" +
			"id SERIAL NOT NULL PRIMARY KEY, " +
			"model varchar(30) NOT NULL, " +
			"company varchar(30) NOT NULL, " +
			"price integer NOT NULL\n);",
	)
	if err != nil {
		panic(err)
	}

	fmt.Print("Table created! Affected: ")
	fmt.Println(result.RowsAffected())
}

func insertProduct(db *sql.DB) int {
	fmt.Println("Inserting data...")

	var id int

	db.QueryRow(
		"insert into Products (model, company, price) values ('Mate 10 Pro', $1, $2) returning id",
		"Huawei",
		35000,
	).Scan(&id)

	fmt.Println("Inserted id:", id)

	return id
}

func getProducts(db *sql.DB) {
	fmt.Println("Receiving data...")

	rows, err := db.Query("select * from Products where price > $1", 1000)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	products := []product{}

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	fmt.Print("Received data: ")
	fmt.Println(products)
}

func updateProduct(db *sql.DB, id int) {
	fmt.Println("Updating data...")

	result, err := db.Exec("update Products set price = $1 where id = $2", 45000, id)
	if err != nil {
		panic(err)
	}

	fmt.Print("Affected: ")
	fmt.Println(result.RowsAffected())
}

func deleteProduct(db *sql.DB, id int) {
	fmt.Println("Deleting data...")

	result, err := db.Exec("delete from Products where id = $1", id)
	if err != nil {
		panic(err)
	}

	fmt.Print("Affected: ")
	fmt.Println(result.RowsAffected())
}
