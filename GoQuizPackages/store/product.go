package store

import "github.com/google/uuid"

type Product struct {
	prodId   uuid.UUID
	name     string
	Category Category
	Stock    int
	Price    float64
}

type ProductList struct {
	prod []Product
}

// constructor
func NewProduct(name string, category Category, stock int, price float64) *Product {
	return &Product{
		prodId:   uuid.New(),
		name:     name,
		Category: category,
		Stock:    stock,
		Price:    price,
	}
}
