package write

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	Uuid       string `gorm:"index;unique"`
	CustomerId string
	PaymentId  string
	CreatedAt  time.Time
}

// IOrderSaver persists order to db, filesystem, etc.
type IOrderSaver interface {
	Save(order *Order) error
}

// OrderSaver implementation.
type OrderSaver struct {
	orm    *gorm.DB
	orders map[string]*Order
}

// NewOrderSaver Constructor.
func NewOrderSaver(orm *gorm.DB) *OrderSaver {
	return &OrderSaver{
		orm:    orm,
		orders: make(map[string]*Order),
	}
}

func (repo *OrderSaver) Save(order *Order) error {
	result := repo.orm.Create(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
