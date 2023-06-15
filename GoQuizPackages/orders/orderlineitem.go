package orders

import (
	"codeid.goquizpackage/store"
	"github.com/google/uuid"
)

type OrderLineItem struct {
	ordlineId uuid.UUID
	Product   store.Product
	qty       int
	subTotal  float64
	Order     Orders
}

func NewOrderLineItem(product store.Product, qty int, subTotal float64, order Orders) *OrderLineItem {
	return &OrderLineItem{
		ordlineId: uuid.New(),
		Product:   product,
		qty:       qty,
		subTotal:  subTotal,
		Order:     order,
	}
}
