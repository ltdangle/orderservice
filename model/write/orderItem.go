package write

import (
	"errors"
	"gorm.io/gorm"
)

type OrderItem struct {
	Uuid        string `gorm:"primaryKey;index;unique"`
	OrderId     string
	ProductId   string
	Title       string
	Description string
	Price       int
}

type IOrderModifier interface {
	AddItem(product *OrderItem) error
	DeleteItem(productId string, orderId string) error
}

type OrderModifier struct {
	orm *gorm.DB
}

func NewOrderModifier(orm *gorm.DB) *OrderModifier {
	return &OrderModifier{orm: orm}
}

func (om *OrderModifier) AddItem(product *OrderItem) error {
	result := om.orm.Create(product)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (om *OrderModifier) DeleteItem(productId string, orderId string) error {
	result := om.orm.Where("uuid = ? AND order_id = ?", productId, orderId).Delete(&OrderItem{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Item not found")
	}

	return nil
}
