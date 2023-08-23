package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"orders/actions"
	"orders/model/read"
)

// Cache interface.
type Cache interface {
	Get(key string, obj any) error
	Set(key string, value any) error
}

// RetrieveOrder controller.
type RetrieveOrder struct {
	action *actions.RetrieveOrder
	cache  Cache
	rspndr *Responder
}

// NewRetrieveOrder constructor.
func NewRetrieveOrder(action *actions.RetrieveOrder, cache Cache, rspndr *Responder) *RetrieveOrder {
	return &RetrieveOrder{
		action: action,
		cache:  cache,
		rspndr: rspndr,
	}
}

func (c *RetrieveOrder) Retrieve(w http.ResponseWriter, r *http.Request) {
	var order *read.Order

	vars := mux.Vars(r)
	uuid := vars["uuid"]

	// Check cache.
	cached := &read.Order{}
	err := c.cache.Get(uuid, cached)

	// Cache exists.
	if err == nil {
		c.rspndr.Success(w, cached)
		return
	}

	// Cache doesn't exit, retrieve order and cache it.
	order, err = c.action.Retrieve(uuid)
	if err != nil {
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Cache order.
	err = c.cache.Set(uuid, order)
	if err != nil {
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.rspndr.Success(w, order)
}
