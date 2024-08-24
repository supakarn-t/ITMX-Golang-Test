package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"ITMX-Golang-Test/models"

	"gorm.io/gorm"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if customer.Name == "" || customer.Age <= 0 {
		http.Error(w, "Invalid customer data", http.StatusBadRequest)
		return
	}

	err = db.Create(&customer).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetCustomer(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	idStr := strings.TrimPrefix(r.URL.Path, "/customers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var customer models.Customer
	err = db.First(&customer, id).Error
	if err != nil {
		http.Error(w, fmt.Sprintf("Not found customer with id %d", id), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	idStr := strings.TrimPrefix(r.URL.Path, "/customers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var customer models.Customer
	err = db.First(&customer, id).Error
	if err != nil {
		http.Error(w, fmt.Sprintf("Not found customer with id %d", id), http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if customer.Name == "" || customer.Age <= 0 {
		http.Error(w, "Invalid customer data", http.StatusBadRequest)
		return
	}
	err = db.Save(&customer).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	idStr := strings.TrimPrefix(r.URL.Path, "/customers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result := db.Delete(&models.Customer{}, id)
	if result.Error != nil || result.RowsAffected == 0 {
		http.Error(w, fmt.Sprintf("Not found customer with id %d", id), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(fmt.Sprintf("Customer with id %d deleted", id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
