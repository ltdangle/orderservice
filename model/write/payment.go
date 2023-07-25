package write

import "time"

type Payment struct {
	Uuid      string `gorm:"index;unique"`
	Date      time.Time
	OrderId   string
	PaymentId string
}
