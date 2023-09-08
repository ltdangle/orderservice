package rest

import (
	"fmt"
	"net/http"
	"orders/actions"
	"orders/infra"

	"github.com/gorilla/mux"
)

// Create order controller.
//
//go:generate newc
type RetrieveOrder struct {
	action *actions.RetrieveOrder
	rspndr *Responder
	logger infra.Logger
}

func (c *RetrieveOrder) Retrieve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	order, err := c.action.Retrieve(uuid)

	if err != nil {
		c.logger.Log("rest/retrieveOrder: " + err.Error())
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.logger.Log(fmt.Sprintf("rest/retrieveOrder: %v", order))
	c.rspndr.Success(w, order)
}
