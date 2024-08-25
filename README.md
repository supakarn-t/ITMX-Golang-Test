# ITMX-Golang-Test

This project is a CRUD REST API written in Go, designed to manage customers. It uses SQLite as the database and the GORM framework for ORM.

## Getting Started

### Prerequisites

- Go 1.16 or later
- SQLite

### Installation

1. Clone the repository:

```sh
git clone https://github.com/supakarn-t/ITMX-Golang-Test.git
cd ITMX-Golang-Test
```

2. Install dependencies:

```sh
go mod tidy
```

3. Set up the database:

```sh
go run main.go
```

## Running the Tests

To run the tests, use the following command:

```sh
go test ./handlers -v
go test -cover ./handlers
```

## API Endpoints

### Create a Customer

- URL: `/customers`
- Method: `POST`
- Request Body:

```sh
{
  "name": "Alice",
  "age": 28
}
```

- Response:

```sh
{
  "id": 1,
  "name": "Alice",
  "age": 28
}
```

### Get a Customer

- URL: `/customers/{id}`
- Method: `GET`
- Response:

```sh
{
  "id": 1,
  "name": "John Doe",
  "age": 30
}
```

### Update a Customer

- URL: `/customers/{id}`
- Method: `PUT`
- Request Body:

```sh
{
  "name": "John Smith",
  "age": 35
}
```

- Response:

```sh
{
  "id": 1,
  "name": "John Smith",
  "age": 35
}
```

### Delete a Customer

- URL: `/customers/{id}`
- Method: `DELETE`
- Response:

```sh
{
  "message": "Customer with id 1 deleted"
}
```
