package repository

import "orders/entity"

// Persists Order to db, filesystem, etc.
type orderRepository interface {
	Save(order entity.Order)
	FindById(orderId string) *entity.Order
}
