package rest

import (
	"encoding/json"
	"net/http"
	"orders/actions"
)

type OrderModifier struct {
	action *actions.OrderActions
	rspndr *Responder
}

func NewOrderModifier(action *actions.OrderActions, rspndr *Responder) *OrderModifier {
	return &OrderModifier{
		action: action,
		rspndr: rspndr,
	}
}

func (c *OrderModifier) AddProduct(w http.ResponseWriter, r *http.Request) {
	var request actions.AddProductRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		c.rspndr.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err, item := c.action.AddProduct(request)

	if err != nil {
		if err.Error() == "order not found" {
			c.rspndr.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.rspndr.Success(w, item)
}

func (c *OrderModifier) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var request actions.DeleteProductRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		c.rspndr.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.action.DeleteProduct(request)

	if err != nil {
		if err.Error() == "item not found" {
			c.rspndr.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.rspndr.Success(w, "")
}
