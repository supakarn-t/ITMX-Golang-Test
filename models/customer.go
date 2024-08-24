package models

type Customer struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
