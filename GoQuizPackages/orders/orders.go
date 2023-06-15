package orders

import (
	"codeid.goquizpackage/users"
	"github.com/google/uuid"
)

type Orders struct {
	orderId    uuid.UUID
	orderNo    string
	User       users.Users
	totalPrice float64
	status     string
}

func NewOrders(orderNo string, user users.Users, totalPrice float64, status string) *Orders {
	return &Orders{
		orderId:    uuid.New(),
		orderNo:    orderNo,
		User:       user,
		totalPrice: totalPrice,
		status:     status,
	}
}
