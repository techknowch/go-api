package handlers

import (
	"net/http"
	"encoding/json"
	"go-api/internal/models"
	"gp-api/internal/services"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	products := services.GetUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func HandlePostProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdProduct = services.CreateProduct(product)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	products := services.GetProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}