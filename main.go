package main

import (
	"ITMX-Golang-Test/database"
	"ITMX-Golang-Test/handlers"
	"log"
	"net/http"
)

func main() {
	db := database.SetupDatabase()

	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			handlers.CreateCustomer(w, r, db)
		default:
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/customers/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetCustomer(w, r, db)
		case "PUT":
			handlers.UpdateCustomer(w, r, db)
		case "DELETE":
			handlers.DeleteCustomer(w, r, db)
		default:
			http.NotFound(w, r)
		}
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
