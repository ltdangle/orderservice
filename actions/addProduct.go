package actions

import (
	"errors"
	"github.com/google/uuid"
	"orders/model/read"
	"orders/model/write"
)

type AddProductRequest struct {
	OrderID string
}

type DeleteProductRequest struct {
	OrderID string
	ItemID  string
}

type ProductAdder struct {
	repo   write.IOrderModifier
	finder read.OrderFinderById
}

func NewProductAdder(repo write.IOrderModifier, finder read.OrderFinderById) *ProductAdder {
	return &ProductAdder{
		repo:   repo,
		finder: finder,
	}
}

func (action *ProductAdder) AddProduct(r AddProductRequest) (error, *write.OrderItem) {
	order, err := action.finder.Find(r.OrderID)

	if order == nil {
		return errors.New("order not found"), nil
	}

	if err != nil {
		return err, nil
	}

	orderItem := write.OrderItem{
		Uuid:        uuid.New().String(),
		OrderId:     r.OrderID,
		ProductId:   uuid.New().String(),
		Title:       "one more item",
		Description: "description",
		Price:       100,
	}

	err = action.repo.AddItem(&orderItem)

	if err != nil {
		return err, nil
	}

	return nil, &orderItem
}
