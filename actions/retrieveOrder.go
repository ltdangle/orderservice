package actions

import (
	"errors"
	"orders/model/read"
)

// RetrieveOrder action.
type RetrieveOrder struct {
	findOrder read.FindOrder
}

// Constructor.
func NewRetrieveOrder(findOrder read.FindOrder) *RetrieveOrder {
	return &RetrieveOrder{
		findOrder: findOrder,
	}
}

func (action *RetrieveOrder) Retrieve(uuid string) (*read.Order, error) {
	order, err := action.findOrder(uuid)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order is nil")
	}
	return order, nil
}
