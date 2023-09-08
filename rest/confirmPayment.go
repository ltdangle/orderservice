package rest

import (
	"net/http"
	"orders/actions"
	"orders/infra"

	"github.com/gorilla/mux"
)

// Create order controller.
//
//go:generate newc
type ConfirmPayment struct {
	confirmPayment *actions.ConfirmPayment
	rspndr         *Responder
	logger         infra.Logger
}

func (a *ConfirmPayment) ConfirmPayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderUuid := vars["uuid"]
	paymentUuid := vars["paymentUuid"]

	err := a.confirmPayment.Confirm(orderUuid, paymentUuid)

	if err != nil {
		a.logger.Log("rest/confirmPayment: " + err.Error())
		a.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Redirect to view order
	a.logger.Log("rest/confirmPayment: redirect to " + "/order/" + orderUuid)
	http.Redirect(w, r, "/order/"+orderUuid, http.StatusSeeOther)
}
