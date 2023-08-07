package rest

import (
	"net/http"
	"orders/actions"
	"orders/infra/cache"

	"github.com/gorilla/mux"
)

// Create order controller.
type RetrieveOrder struct {
	action *actions.RetrieveOrder
	cache  cache.Cache
	rspndr *Responder
}

// Constructor.
func NewRetrieveOrder(action *actions.RetrieveOrder, cache cache.Cache, rspndr *Responder) *RetrieveOrder {
	return &RetrieveOrder{
		action: action,
		cache:  cache,
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
