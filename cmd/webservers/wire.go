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

func app(orderSaver write.IOrderSaver, sqlDb *sql.DB, dateFormat string, url actions.CheckoutUrl, db *gorm.DB, orderModifier write.IOrderModifier) *webApp {
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
		//read.NewOrderItemFinderByItemId,
		read.NewOrderItemFinderById,
	)
	return &webApp{}
}
