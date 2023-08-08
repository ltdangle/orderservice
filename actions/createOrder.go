package actions

import (
	"github.com/google/uuid"
	"orders/model/write"
	"time"
)

// NewOrderRequest.
type NewOrderRequest struct{}

// CreateOrder action.
type CreateOrder struct {
	repo write.IOrderSaver
}

// Constructor.
func NewCreateOrder(repo write.IOrderSaver) *CreateOrder {
	return &CreateOrder{
		repo: repo,
	}
}

func (action *CreateOrder) Create(r NewOrderRequest) (*write.Order, error) {
	order := &write.Order{
		Uuid:      uuid.New().String(),
		CreatedAt: time.Now(),
	}

	err := action.repo.Save(order)

	if err != nil {
		return nil, err
	}

	return order, nil
}
