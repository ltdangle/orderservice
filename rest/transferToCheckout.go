package rest

import (
	"net/http"
	"orders/actions"
	"orders/infra"

	"github.com/gorilla/mux"
)

//go:generate newc
type CheckoutTransfer struct {
	retrieve *actions.RetrieveOrder
	checkout *actions.CheckoutTransfer
	rspndr   *Responder
	logger   infra.Logger
}

func (c *CheckoutTransfer) Checkout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	// Retrieve order.
	order, err := c.retrieve.Retrieve(uuid)

	if err != nil {
		c.logger.Log("rest/transferToCheckout: " + err.Error())
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Create checkout url and return it.
	type payload struct {
		Url string
	}
	checkoutUrl := c.checkout.Url(order)

	c.logger.Log("rest/transferToCheckout: transfer to checkout url " + string(checkoutUrl))
	c.rspndr.Success(w, payload{Url: checkoutUrl})

}
