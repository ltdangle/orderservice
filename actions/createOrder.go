package actions

import (
	"github.com/google/uuid"
	"orders/model/write"
	"time"
)

// NewOrderRequest.
type NewOrderRequest struct{}

// Save order function.
type SaveOrder func(order *write.Order) error

// CreateOrder action.
type CreateOrder struct {
	saveOrder SaveOrder
}

// Constructor.
func NewCreateOrder(saveOrder SaveOrder) *CreateOrder {
	return &CreateOrder{
		saveOrder: saveOrder,
	}
}

func (action *CreateOrder) Create(r NewOrderRequest) (*write.Order, error) {
	order := &write.Order{
		Uuid:      uuid.New().String(),
		Status:    "created",
		CreatedAt: time.Now(),
	}

	err := action.saveOrder(order)

	if err != nil {
		return nil, err
	}

	return order, nil
}
