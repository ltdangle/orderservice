package entity

import "time"

type Order struct {
	Uuid       string
	CustomerId string
	Products   []Product
	PaymentId  string
	CreatedAt  time.Time
}
