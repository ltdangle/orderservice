package rest

import (
	"net/http"
	"orders/actions"
)

// Create order controller.
type CreateOrder struct {
	action *actions.CreateOrder
	rspndr *Responder
}

// Constructor.
func NewCreateOrder(action *actions.CreateOrder, rspndr *Responder) *CreateOrder {
	return &CreateOrder{
		action: action,
		rspndr: rspndr,
	}
}

func (c *CreateOrder) Create(w http.ResponseWriter, r *http.Request) {
	order, err := c.action.Create(actions.NewOrderRequest{})

	if err != nil {
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
	}

	c.rspndr.Success(w, order)
}
