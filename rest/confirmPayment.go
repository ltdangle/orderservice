package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"orders/actions"
)

// Create order controller.
type ConfirmPayment struct {
	confirmPayment *actions.ConfirmPayment
	rspndr         *Responder
}

func NewConfirmPayment(confirmPayment *actions.ConfirmPayment, rspndr *Responder) *ConfirmPayment {
	return &ConfirmPayment{confirmPayment: confirmPayment, rspndr: rspndr}
}

func (a *ConfirmPayment) ConfirmPayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderUuid := vars["uuid"]
	paymentUuid := vars["paymentUuid"]

	err := a.confirmPayment.Confirm(orderUuid, paymentUuid)
	if err != nil {
		return
	}

	if err != nil {
		a.rspndr.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Redirect to view order
	http.Redirect(w, r, "/order/"+orderUuid, http.StatusSeeOther)

}
