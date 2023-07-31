package actions

import "orders/entity"

// Action interfaces to be implemented.
type ProductAdder interface {
	AddProduct(order *entity.Order, product *entity.Product)
}

type ProductDeleter interface {
	DeleteProduct(order *entity.Order, productId string)
}

type ProductConfirmer interface {
	ConfirmProduct(productId string) bool
}

type OrderCalculator interface {
	Total(order *entity.Order)
}

type PaymentConfirmer interface {
	Confirm(payment *entity.Payment)
}
