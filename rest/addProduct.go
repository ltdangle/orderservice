package rest

import (
	"encoding/json"
	"net/http"
	"orders/actions"
)

type AddProduct struct {
	action *actions.ProductAdder
	rspndr *Responder
}

func NewAddProduct(action *actions.ProductAdder, rspndr *Responder) *AddProduct {
	return &AddProduct{
		action: action,
		rspndr: rspndr,
	}
}

func (c *AddProduct) AddProduct(w http.ResponseWriter, r *http.Request) {
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
