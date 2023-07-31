package main

import (
	"log"
	"net/http"
	"orders/actions"
	"orders/repository"
	"orders/rest"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Entity repository.
	repo := repository.NewInMemoryOrderRepo()

	// Rest json responder.
	respndr:=rest.NewResponder("2006-01-02 15:04:05")

	// New order controller.
	cntrlr := rest.NewCreateOrder(actions.NewCreateOrder(repo), respndr)

	// Router and server.
	r := mux.NewRouter()
	r.HandleFunc("/create", cntrlr.Create)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
