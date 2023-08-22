package rest

import (
	"net/http"
	"orders/actions"
	"orders/helpers"

	"github.com/gorilla/mux"
)

// Create order controller.
type CheckoutTransfer struct {
	retrieve *actions.RetrieveOrder
	checkout *actions.CheckoutTransfer
	rspndr   *Responder
}

// Constructor.
func NewCheckoutTransfer(checkout *actions.CheckoutTransfer, retrieve *actions.RetrieveOrder, rspndr *Responder) *CheckoutTransfer {
	return &CheckoutTransfer{
		retrieve: retrieve,
		checkout: checkout,
		rspndr:   rspndr,
	}
}

func (c *CheckoutTransfer) Checkout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	// Retrieve order.
	order, err := c.retrieve.Retrieve(uuid)

	if err != nil {
		c.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Create checkout url and return it.
	type payload struct {
		Url string
	}
	checkoutUrl := c.checkout.Url(helpers.NullStringToString(order.Uuid))

	c.rspndr.Success(w, payload{Url: checkoutUrl})

}