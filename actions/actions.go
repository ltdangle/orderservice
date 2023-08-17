package actions

import (
	"errors"
	"github.com/google/uuid"
	"orders/model/read"
	"orders/model/write"
)

// Action interfaces to be implemented.
type ProductAdder interface {
	AddProduct(order *write.Order, product *write.OrderItem)
}

type AddProductRequest struct {
	OrderID string
}

type ProductDeleter interface {
	DeleteProduct(order *write.Order, productId string)
}

type DeleteProductRequest struct {
	OrderID string
	ItemID  string
}

type ProductConfirmer interface {
	ConfirmProduct(productId string) bool
}

type OrderCalculator interface {
	Total(order *write.Order)
}

type PaymentConfirmer interface {
	Confirm(payment *write.Payment)
}

type OrderActions struct {
	repo   write.IOrderModifier
	finder read.OrderFinderById
}

// Constructor.
func NewOrderActions(repo write.IOrderModifier, finder read.OrderFinderById) *OrderActions {
	return &OrderActions{
		repo:   repo,
		finder: finder,
	}
}

func (action *OrderActions) AddProduct(r AddProductRequest) (error, *write.OrderItem) {
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

func (action *OrderActions) DeleteProduct(r DeleteProductRequest) error {

	return action.repo.DeleteItem(r.ItemID, r.OrderID)
}
