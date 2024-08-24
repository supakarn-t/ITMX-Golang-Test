package database

import (
	"log"

	"ITMX-Golang-Test/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("customers.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	db.AutoMigrate(&models.Customer{})
	SeedDatabase(db)
	return db
}

func SeedDatabase(db *gorm.DB) {
	db.Create(&models.Customer{Name: "John Doe", Age: 30})
	db.Create(&models.Customer{Name: "Jane Doe", Age: 25})
}
