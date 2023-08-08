package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"orders/actions"
)

// Create order controller.
type RetrieveOrder struct {
	action *actions.RetrieveOrder
	rspndr *Responder
}

// Constructor.
func NewRetrieveOrder(action *actions.RetrieveOrder, rspndr *Responder) *RetrieveOrder {
	return &RetrieveOrder{
		action: action,
		rspndr: rspndr,
	}
}

func (c *RetrieveOrder) Retrieve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	order, err := c.action.Retrieve(uuid)

	if err != nil {
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
	}

	c.rspndr.Success(w, order)
}
