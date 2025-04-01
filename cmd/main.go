package main

import (
	"net/http"
	"github/gin-gonic/gin"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

var products = []Product{
	{ID: "1", Name: "Product 1", Price: 10.0, Stock: 100},
	{ID: "2", Name: "Product 2", Price: 20.0, Stock: 200},
	{ID: "3", Name: "Product 3", Price: 30.0, Stock: 300},
	{ID: "4", Name: "Product 4", Price: 40.0, Stock: 400},
	{ID: "5", Name: "Product 5", Price: 50.0, Stock: 500},
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)
	router.Run("localhost:8080")
	}