//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"orders/actions"
	"orders/model/read"
	"orders/model/write"
	"orders/rest"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func buildDependencies(sqlDb *sql.DB, db *gorm.DB, dateFormat string, url actions.CheckoutUrl) *webApp {
	wire.Build(
		newWebApp,
		rest.NewCreateOrder,
		rest.NewAddProduct,
		rest.NewDeleteProduct,
		rest.NewRetrieveOrder,
		rest.NewCheckoutTransfer,
		rest.NewConfirmPayment,
		rest.NewResponder,
		actions.NewCreateOrder,
		actions.NewProductAdder,
		actions.NewProductDeleter,
		actions.NewConfirmPayment,
		actions.NewRetrieveOrder,
		actions.NewCheckoutTransfer,
		read.NewOrderFinderById,
		read.NewOrderFinderActiveById,
		read.NewOrderItemFinderById,
		write.NewOrderSaver,
		write.NewOrderModifier,
	)
	return &webApp{}
}
