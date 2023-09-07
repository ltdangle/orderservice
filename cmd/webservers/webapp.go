package main

import (
	"log"
	"net/http"
	"orders/rest"
	"time"

	"github.com/gorilla/mux"
)

type webApp struct {
	createOrderCntrlr   *rest.CreateOrder
	retrieveOrderCntrlr *rest.RetrieveOrder
	checkoutCntrlr      *rest.CheckoutTransfer
	confirmPmntCntrlr   *rest.ConfirmPayment
	addProductCntrl     *rest.AddProduct
	modifyOrderCntrlr   *rest.DeleteProduct
}

func newWebApp(
	createOrderCntrlr *rest.CreateOrder,
	retrieveOrderCntrlr *rest.RetrieveOrder,
	checkoutCntrlr *rest.CheckoutTransfer,
	confirmPmntCntrlr *rest.ConfirmPayment,
	addProductCntrl *rest.AddProduct,
	modifyOrderCntrlr *rest.DeleteProduct,
) *webApp {
	return &webApp{createOrderCntrlr: createOrderCntrlr, retrieveOrderCntrlr: retrieveOrderCntrlr, checkoutCntrlr: checkoutCntrlr, confirmPmntCntrlr: confirmPmntCntrlr, addProductCntrl: addProductCntrl, modifyOrderCntrlr: modifyOrderCntrlr}
}

func (app *webApp) run() {
	r := mux.NewRouter()
	r.HandleFunc("/order/create", app.createOrderCntrlr.Create).Methods("GET")
	r.HandleFunc("/order/{uuid}", app.retrieveOrderCntrlr.Retrieve).Methods("GET")
	r.HandleFunc("/order/{uuid}/checkout", app.checkoutCntrlr.Checkout).Methods("GET")
	r.HandleFunc("/order/{uuid}/payment/{paymentUuid}", app.confirmPmntCntrlr.ConfirmPayment).Methods("GET")
	r.HandleFunc("/product/add", app.addProductCntrl.AddProduct).Methods("POST")
	r.HandleFunc("/product/delete", app.modifyOrderCntrlr.DeleteProduct).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
