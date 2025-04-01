package services

import (
	"errors"
	"go-api/internal/models"
)

var users []models.User
var products []models.Product

func GetUsers() ([]models.User, error) {
	if len(users) == 0 {
		return nil, errors.New("no users found")
	}
	return users, nil
}

func CreateProduct(product models.Product) (models.Product, error) {
	if product.ID == "" || product.Name == "" {
		return models.Product{}, errors.New("invalid product data")
	}
	products = append(products, product)
	return product, nil
}

func GetProducts() ([]models.Product, error) {
	if len(products) == 0 {
		return nil, errors.New("no products found")
	}
	return products, nil
}