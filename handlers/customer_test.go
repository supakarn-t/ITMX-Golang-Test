package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"ITMX-Golang-Test/database"
	"ITMX-Golang-Test/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	err := os.Remove("customers.db")
	if err != nil {
		log.Fatal(err)
	}
	db = database.SetupDatabase()
}

func TestCreateCustomer(t *testing.T) {
	customer := models.Customer{Name: "Alice", Age: 28}

	payload, err := json.Marshal(customer)
	require.NoError(t, err, "Failed to marshal customer payload")

	req, err := http.NewRequest("POST", "/customers", bytes.NewBuffer(payload))
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CreateCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusCreated, recorder.Code, "Expected status code 201")

	var createdCustomer models.Customer

	err = json.Unmarshal(recorder.Body.Bytes(), &createdCustomer)
	require.NoError(t, err, "Failed to unmarshal response body")

	assert.NotZero(t, createdCustomer.ID, "Customer ID should not be zero")
	assert.Equal(t, customer.Name, createdCustomer.Name, "Customer name mismatch")
	assert.Equal(t, customer.Age, createdCustomer.Age, "Customer age mismatch")
}

func TestCreateCustomer_InvalidName(t *testing.T) {
	customer := models.Customer{Name: "", Age: 30}

	payload, err := json.Marshal(customer)
	require.NoError(t, err, "Failed to marshal customer payload")

	req, err := http.NewRequest("POST", "/customers", bytes.NewBuffer(payload))
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CreateCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusBadRequest, recorder.Code, "Expected status code 400")
}

func TestCreateCustomer_InvalidAge(t *testing.T) {
	customer := models.Customer{Name: "Alice"}

	payload, err := json.Marshal(customer)
	require.NoError(t, err, "Failed to marshal customer payload")

	req, err := http.NewRequest("POST", "/customers", bytes.NewBuffer(payload))
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CreateCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusBadRequest, recorder.Code, "Expected status code 400")
}

func TestCreateCustomer_MissingData(t *testing.T) {
	customer := models.Customer{Age: 28}

	payload, err := json.Marshal(customer)
	require.NoError(t, err, "Failed to marshal customer payload")

	req, err := http.NewRequest("POST", "/customers", bytes.NewBuffer(payload))
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CreateCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusBadRequest, recorder.Code, "Expected status code 400")
}

func TestGetCustomer(t *testing.T) {
	req, err := http.NewRequest("GET", "/customers/1", nil)
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code, "Expected status code 200")

	var customer models.Customer
	err = json.Unmarshal(recorder.Body.Bytes(), &customer)
	require.NoError(t, err, "Failed to unmarshal response body")

	assert.Equal(t, uint(1), customer.ID, "Customer ID mismatch")
	assert.Equal(t, "John Doe", customer.Name, "Customer name mismatch")
	assert.Equal(t, 30, customer.Age, "Customer age mismatch")
}

func TestGetCustomer_NotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/customers/999", nil)
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusNotFound, recorder.Code, "Expected status code 404")
}

func TestGetCustomer_InvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/customers/invalidID", nil)
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusBadRequest, recorder.Code, "Expected status code 400")
}

func TestUpdateCustomer(t *testing.T) {
	customer := models.Customer{ID: 1, Name: "John Smith", Age: 35}
	payload, err := json.Marshal(customer)
	require.NoError(t, err, "Failed to marshal customer payload")

	req, err := http.NewRequest("PUT", "/customers/1", bytes.NewBuffer(payload))
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UpdateCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code, "Expected status code 200")

	var updatedCustomer models.Customer

	err = json.Unmarshal(recorder.Body.Bytes(), &updatedCustomer)
	require.NoError(t, err, "Failed to unmarshal response body")

	assert.Equal(t, uint(1), updatedCustomer.ID, "Customer ID mismatch")
	assert.Equal(t, customer.Name, updatedCustomer.Name, "Customer name mismatch")
	assert.Equal(t, customer.Age, updatedCustomer.Age, "Customer age mismatch")
}

func TestUpdateCustomer_NotFound(t *testing.T) {
	req, err := http.NewRequest("PUT", "/customers/999", nil)
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UpdateCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusNotFound, recorder.Code, "Expected status code 404")
}

func TestUpdateCustomer_InvalidID(t *testing.T) {
	req, err := http.NewRequest("PUT", "/customers/invalidID", nil)
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UpdateCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusBadRequest, recorder.Code, "Expected status code 400")
}

func TestUpdateCustomer_InvalidData(t *testing.T) {
	customer := models.Customer{ID: 1, Name: "", Age: -1}

	payload, err := json.Marshal(customer)
	require.NoError(t, err, "Failed to marshal customer payload")

	req, err := http.NewRequest("PUT", "/customers/1", bytes.NewBuffer(payload))
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UpdateCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusBadRequest, recorder.Code, "Expected status code 400")
}

func TestDeleteCustomer(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/customers/1", nil)
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		DeleteCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code, "Expected status code 200")

	var response string

	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	require.NoError(t, err, "Failed to unmarshal response body")

	assert.Contains(t, response, "deleted")

	var customer models.Customer
	result := db.First(&customer, 1)

	assert.Error(t, result.Error, "Expected error when fetching deleted customer")
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error, "Expected record not found error")
}

func TestDeleteCustomer_NotFound(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/customers/999", nil)
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		DeleteCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusNotFound, recorder.Code, "Expected status code 404")
}

func TestDeleteCustomer_InvalidID(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/customers/invalidID", nil)
	require.NoError(t, err, "Failed to create new request")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		DeleteCustomer(w, r, db)
	})

	handler.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusBadRequest, recorder.Code, "Expected status code 400")
}
