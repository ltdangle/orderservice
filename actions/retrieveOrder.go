package actions

import (
	"orders/model/read"
)

// RetrieveOrder action.
type RetrieveOrder struct {
	finder read.OrderFinderById
}

// Constructor.
func NewRetrieveOrder(finder read.OrderFinderById) *RetrieveOrder {
	return &RetrieveOrder{
		finder: finder,
	}
}

func (action *RetrieveOrder) Retrieve(uuid string) (*read.Order, error) {
	order, err := action.finder.Find(uuid)
	if err != nil {
		return nil, err
	}

	return order, nil
}
