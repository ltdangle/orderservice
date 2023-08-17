package rest

import (
	"encoding/json"
	"net/http"
	"orders/actions"
)

type DeleteProduct struct {
	action *actions.ProductDeleter
	rspndr *Responder
}

// Constructor.
func NewDeleteProduct(action *actions.ProductDeleter, rspndr *Responder) *DeleteProduct {
	return &DeleteProduct{
		action: action,
		rspndr: rspndr,
	}
}

func (c *DeleteProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
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
