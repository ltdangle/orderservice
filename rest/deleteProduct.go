package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orders/actions"
	"orders/infra"
)

//go:generate newc
type DeleteProduct struct {
	action *actions.ProductDeleter
	rspndr *Responder
	logger infra.Logger
}

func (c *DeleteProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var request actions.DeleteProductRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		c.logger.Log("rest/deleteProduct: " + err.Error())
		c.rspndr.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.action.DeleteProduct(request)

	if err != nil {
		if err.Error() == "item not found" {
			c.logger.Log("rest/deleteProduct: " + err.Error())
			c.rspndr.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		c.logger.Log("rest/deleteProduct: " + err.Error())
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.logger.Log(fmt.Sprintf("rest/deleteProduct: %v", request))
	c.rspndr.Success(w, "")
}
