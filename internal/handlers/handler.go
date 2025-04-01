package handlers

import (
	"net/http"
	"encoding/json"
	"go-api/internal/models"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	products := models.GetUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func HandlePostProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.CreateProduct(product)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}