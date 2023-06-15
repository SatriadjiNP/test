package orders

import (
	"codeid.goquizpackage/store"
	"codeid.goquizpackage/users"
	"github.com/google/uuid"
)

type Cart []ItemProduct

func AddToCart(Product *store.Product, qty int, User *users.Users) []ItemProduct {
	subTotal := float64(qty) * Product.Price
	return []ItemProduct{
		{uuid.New(), *Product, qty, subTotal, *User},
	}
}
