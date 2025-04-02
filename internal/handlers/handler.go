package handlers

import (
    "github.com/gin-gonic/gin"
    "go-api/internal/models"
    "go-api/internal/services"
    "net/http"
)

func HandleGetUser(c *gin.Context) {
    users, err := services.GetUsers()
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func HandlePostProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    createdProduct, err := services.CreateProduct(product)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, createdProduct)
}

func HandleGetProducts(c *gin.Context) {
    products, err := services.GetProducts()
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, products)
}

func HandleDeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func HandleGetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := services.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}