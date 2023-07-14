package entity

import "time"

type Order struct {
	customerId string
	products   []Product
	paymentId  string
	createdAt  time.Time
}
