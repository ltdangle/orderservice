package actions

import (
	"github.com/google/uuid"
	"orders/entity"
	"orders/repository"
	"time"
)

// NewOrderRequest.
type NewOrderRequest struct{}

// CreateOrder action.
type CreateOrder struct {
	repo repository.OrderRepository
}

// Constructor.
func NewCreateOrder(repo repository.OrderRepository) *CreateOrder {
	return &CreateOrder{
		repo: repo,
	}
}

func (action *CreateOrder) Create(r NewOrderRequest) (*entity.Order, error) {
	order := &entity.Order{
		Uuid:      uuid.New().String(),
		CreatedAt: time.Now(),
	}

	err := action.repo.Save(order)

	if err != nil {
		return nil, err
	}

	return order, nil
}
