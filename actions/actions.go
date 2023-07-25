package actions

import (
	"orders/model/write"
)

// Action interfaces to be implemented.
type ProductAdder interface {
	AddProduct(order *write.Order, product *write.OrderItem)
}

type ProductDeleter interface {
	DeleteProduct(order *write.Order, productId string)
}

type ProductConfirmer interface {
	ConfirmProduct(productId string) bool
}

type OrderCalculator interface {
	Total(order *write.Order)
}

type PaymentConfirmer interface {
	Confirm(payment *write.Payment)
}
