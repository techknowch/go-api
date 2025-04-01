package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/internal/handlers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/users", handlers.HandleGetUser)
	router.POST("/products", handlers.HandlePostProduct)
	router.GET("/products", handlers.HandleGetProducts)
}