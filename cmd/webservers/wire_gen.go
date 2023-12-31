// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"gorm.io/gorm"
	"orders/actions"
	"orders/infra"
	"orders/model/read"
	"orders/model/write"
	"orders/rest"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func buildDependencies(sqlDb *sql.DB, db *gorm.DB, dateFormat string, url actions.CheckoutUrl) *webApp {
	iOrderSaver := write.NewOrderSaver(db)
	createOrder := actions.NewCreateOrder(iOrderSaver)
	responder := rest.NewResponder(dateFormat)
	logger := infra.NewSimpleLogger()
	restCreateOrder := rest.NewCreateOrder(createOrder, responder, logger)
	iOrderItemFinderByOrderId := read.NewOrderItemFinderById(sqlDb)
	orderFinderById := read.NewOrderFinderById(sqlDb, iOrderItemFinderByOrderId)
	retrieveOrder := actions.NewRetrieveOrder(orderFinderById)
	restRetrieveOrder := rest.NewRetrieveOrder(retrieveOrder, responder, logger)
	checkoutTransfer := actions.NewCheckoutTransfer(url)
	restCheckoutTransfer := rest.NewCheckoutTransfer(retrieveOrder, checkoutTransfer, responder, logger)
	confirmPayment := actions.NewConfirmPayment(iOrderSaver, db)
	restConfirmPayment := rest.NewConfirmPayment(confirmPayment, responder, logger)
	iOrderModifier := write.NewOrderModifier(db)
	orderFinderActiveById := read.NewOrderFinderActiveById(sqlDb)
	productAdder := actions.NewProductAdder(iOrderModifier, orderFinderActiveById)
	addProduct := rest.NewAddProduct(productAdder, responder, logger)
	productDeleter := actions.NewProductDeleter(iOrderModifier, orderFinderById)
	deleteProduct := rest.NewDeleteProduct(productDeleter, responder, logger)
	mainWebApp := newWebApp(restCreateOrder, restRetrieveOrder, restCheckoutTransfer, restConfirmPayment, addProduct, deleteProduct)
	return mainWebApp
}
