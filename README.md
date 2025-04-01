# Go REST API

This project is a simple REST API built with Go. It provides endpoints for managing users and products.

## Project Structure

```
go-rest-api
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handlers         # HTTP request handlers
│   │   └── handler.go
│   ├── models           # Data structures
│   │   └── model.go
│   ├── routes           # API routes
│   │   └── routes.go
│   └── services         # Business logic
│       └── service.go
├── pkg
│   └── utils            # Utility functions
│       └── utils.go
├── go.mod               # Module definition
├── go.sum               # Module checksums
└── README.md            # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd go-rest-api
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/main.go
   ```

## API Usage

### Endpoints

- `GET /users`: Retrieve a list of users.
- `POST /users`: Create a new user.
- `GET /products`: Retrieve a list of products.
- `POST /products`: Create a new product.

### Example Requests

- **Get Users:**
  ```
  curl -X GET http://localhost:8080/users
  ```

- **Create User:**
  ```
  curl -X POST http://localhost:8080/users -d '{"name": "John Doe", "email": "john@example.com"}'
  ```

## License

This project is licensed under the MIT License.