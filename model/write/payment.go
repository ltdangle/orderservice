package write

import "time"

type Payment struct {
	Uuid      string `gorm:"primaryKey;index;unique"`
	Date      time.Time
	OrderId   string
	PaymentId string
}
