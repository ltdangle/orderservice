package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"orders/actions"
	"orders/model/read"
	"orders/model/write"
	"orders/rest"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlDsn := os.Getenv("MYSQL_DSN") // Replace "KEY" with your variable's key
	if mysqlDsn == "" {
		log.Fatal("MYSQL_DSN env variable not set.")
	}

	// ORM.
	orm, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	repo := write.NewOrderSaver(orm)

	// Rest json responder.
	respndr := rest.NewResponder("2006-01-02 15:04:05")

	// Mysql orm connection.
	mysqlDb, errr := sql.Open("mysql", mysqlDsn)
	if errr != nil {
		log.Fatal(errr.Error())
	}

	// Model read/write implementations.
	orderModifier := write.NewOrderModifier(orm)
	orderItemFinder := read.NewOrderFinderById(mysqlDb, read.NewOrderItemFinderById(mysqlDb))
	orderActiveFinder := read.NewOrderFinderActiveById(mysqlDb)

	// Actions.
	retrieveOrderAction := actions.NewRetrieveOrder(read.NewOrderFinderById(mysqlDb, read.NewOrderItemFinderById(mysqlDb)))

	// Controllers.
	createOrderCntrlr := rest.NewCreateOrder(actions.NewCreateOrder(repo), respndr)
	retrieveOrderCntrlr := rest.NewRetrieveOrder(retrieveOrderAction, respndr)
	modifyOrderCntrlr := rest.NewDeleteProduct(actions.NewProductDeleter(orderModifier, orderItemFinder), respndr)
	addProductCntrl := rest.NewAddProduct(actions.NewProductAdder(orderModifier, orderActiveFinder), respndr)
	checkoutCntrlr := rest.NewCheckoutTransfer(actions.NewCheckoutTransfer("http://checkout.url"), retrieveOrderAction, respndr)

	// Router and server.
	r := mux.NewRouter()
	r.HandleFunc("/order/create", createOrderCntrlr.Create).Methods("GET")
	r.HandleFunc("/order/{uuid}", retrieveOrderCntrlr.Retrieve).Methods("GET")
	r.HandleFunc("/add-product", addProductCntrl.AddProduct).Methods("POST")
	r.HandleFunc("/product", modifyOrderCntrlr.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/order/{uuid}/checkout", checkoutCntrlr.Checkout).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
