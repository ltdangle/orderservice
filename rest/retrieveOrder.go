package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"orders/actions"
	"orders/model/read"
)

// ICache interface.
type ICache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

// ISerializer interface.
type ISerializer interface {
	EncodeToString(v interface{}) (string, error)
	DecodeFromString(s string, v interface{}) error
}

// RetrieveOrder controller.
type RetrieveOrder struct {
	action     *actions.RetrieveOrder
	serializer ISerializer
	cache      ICache
	rspndr     *Responder
}

// NewRetrieveOrder constructor.
func NewRetrieveOrder(action *actions.RetrieveOrder, serializer ISerializer, cache ICache, rspndr *Responder) *RetrieveOrder {
	return &RetrieveOrder{
		action:     action,
		serializer: serializer,
		cache:      cache,
		rspndr:     rspndr,
	}
}

func (c *RetrieveOrder) Retrieve(w http.ResponseWriter, r *http.Request) {
	var order *read.Order

	vars := mux.Vars(r)
	uuid := vars["uuid"]

	// Check cache.
	cachedString, err := c.cache.Get(uuid)
	// Cache exists.
	if err == nil {
		err = c.serializer.DecodeFromString(cachedString, &order)
		if err != nil {
			c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		c.rspndr.Success(w, order)
		return
	}

	// Cache doesn't exit, retrieve order and cache it.
	order, err = c.action.Retrieve(uuid)
	if err != nil {
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Cache order.
	encodedOrder, _ := c.serializer.EncodeToString(order)
	err = c.cache.Set(uuid, encodedOrder)
	if err != nil {
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	c.rspndr.Success(w, order)
}
