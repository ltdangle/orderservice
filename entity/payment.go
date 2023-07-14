package entity

import "time"


type Payment struct {
  date time.Time
  orderId string
  paymentId string
}
