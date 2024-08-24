package main

import (
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db := SetupDatabase()

	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			CreateCustomer(w, r, db)
		default:
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/customers/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			GetCustomer(w, r, db)
		case "PUT":
			UpdateCustomer(w, r, db)
		case "DELETE":
			DeleteCustomer(w, r, db)
		default:
			http.NotFound(w, r)
		}
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Customer struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("customers.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	db.AutoMigrate(&Customer{})
	SeedDatabase(db)
	return db
}

func SeedDatabase(db *gorm.DB) {
	db.Create(&Customer{Name: "John Doe", Age: 30})
	db.Create(&Customer{Name: "Jane Doe", Age: 25})
}

func CreateCustomer(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
}

func GetCustomer(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
}
