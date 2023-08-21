package actions

import (
	"encoding/json"
	"errors"
	"fmt"
	"orders/clients"
	"orders/model/read"
	"orders/model/write"
)

type AddProductRequest struct {
	OrderID   string
	ProductID string
}

type ProductAdder struct {
	repo   write.IOrderModifier
	finder read.OrderFinderActiveById
}

func NewProductAdder(repo write.IOrderModifier, finder read.OrderFinderActiveById) *ProductAdder {
	return &ProductAdder{
		repo:   repo,
		finder: finder,
	}
}

func (action *ProductAdder) AddProduct(r AddProductRequest) (error, *write.OrderItem) {
	order, err := action.finder.FindActive(r.OrderID)

	if order == nil {
		return errors.New("order not found"), nil
	}

	if err != nil {
		return err, nil
	}

	err, product := clients.GetProductById(r.ProductID)
	if err != nil {
		return err, nil
	}

	fmt.Println(product)

	orderItem := write.OrderItem{}

	err = json.Unmarshal([]byte(product), &orderItem)
	if err != nil {
		return err, nil
	}

	orderItem.OrderId = r.OrderID
	orderItem.ProductId = r.ProductID
	orderItem.Uuid = "47rg4yut"

	err = action.repo.AddItem(&orderItem)

	if err != nil {
		return err, nil
	}

	return nil, &orderItem
}
