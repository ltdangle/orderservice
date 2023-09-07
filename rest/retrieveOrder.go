package rest

import (
	"net/http"
	"orders/model/read"

	"github.com/gorilla/mux"
)

// Create order controller.
type RetrieveOrder struct {
	findOrder read.FindOrder
	rspndr    *Responder
}

// Constructor.
func NewRetrieveOrder(findOrder read.FindOrder, rspndr *Responder) *RetrieveOrder {
	return &RetrieveOrder{
		findOrder: findOrder,
		rspndr:    rspndr,
	}
}

func (c *RetrieveOrder) Retrieve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	order, err := c.findOrder(uuid)

	if err != nil {
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if order == nil {
		c.rspndr.Error(w, http.StatusInternalServerError, "order is nil")
		return
	}

	c.rspndr.Success(w, order)
}
