package rest

import (
	"fmt"
	"net/http"
	"orders/actions"
	"orders/infra"
)

// Create order controller.
//go:generate newc
type CreateOrder struct {
	action *actions.CreateOrder
	rspndr *Responder
	logger infra.Logger
}

func (c *CreateOrder) Create(w http.ResponseWriter, r *http.Request) {
	order, err := c.action.Create(actions.NewOrderRequest{})

	if err != nil {
		c.logger.Log("rest/createOrder.go: "+err.Error())
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
	}

	c.logger.Log(fmt.Sprintf("rest/createOrder.go: order created %v",order))
	c.rspndr.Success(w, order)
}
