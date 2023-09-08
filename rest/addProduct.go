package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orders/actions"
	"orders/infra"
)

//go:generate newc
type AddProduct struct {
	action *actions.ProductAdder
	rspndr *Responder
	logger infra.Logger
}

func (c *AddProduct) AddProduct(w http.ResponseWriter, r *http.Request) {
	var request actions.AddProductRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		c.rspndr.Error(w, http.StatusBadRequest, err.Error())
		c.logger.Log("rest/addProduct: " + err.Error())
		return
	}

	err, item := c.action.AddProduct(request)

	if err != nil {
		if err.Error() == "order not found" {
			c.rspndr.Error(w, http.StatusBadRequest, err.Error())
			c.logger.Log("rest/addProduct: " + err.Error())
			return
		}
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		c.logger.Log("rest/addProduct: " + err.Error())
		return
	}

	c.rspndr.Success(w, item)
	c.logger.Log(fmt.Sprintf("rest/addProduct: %v", request))
}
