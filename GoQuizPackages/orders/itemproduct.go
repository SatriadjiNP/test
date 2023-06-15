package orders

import (
	"codeid.goquizpackage/store"
	"codeid.goquizpackage/users"
	"github.com/google/uuid"
)

type ItemProduct struct {
	cartId   uuid.UUID
	Product  store.Product
	qty      int
	subtotal float64
	User     users.Users
}

// constructor
func NewItemProduct(product store.Product, qty int, subtotal float64, user users.Users) *ItemProduct {
	return &ItemProduct{
		cartId:   uuid.New(),
		Product:  product,
		qty:      qty,
		subtotal: subtotal,
		User:     user,
	}
}
