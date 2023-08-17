package actions

import (
	"orders/model/read"
	"orders/model/write"
)

type ProductDeleter struct {
	repo   write.IOrderModifier
	finder read.OrderFinderById
}

// Constructor.
func NewProductDeleter(repo write.IOrderModifier, finder read.OrderFinderById) *ProductDeleter {
	return &ProductDeleter{
		repo:   repo,
		finder: finder,
	}
}

func (action *ProductDeleter) DeleteProduct(r DeleteProductRequest) error {
	return action.repo.DeleteItem(r.ItemID, r.OrderID)
}
