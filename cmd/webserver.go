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
	"orders/infra/cache"
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
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	repo := write.NewOrderSaver(db)

	// Rest json responder.
	respndr := rest.NewResponder("2006-01-02 15:04:05")

	// Mysql db connection.
	mysqlDb, errr := sql.Open("mysql", mysqlDsn)
	if errr != nil {
		log.Fatal(errr.Error())
	}

	// Cache.
	c := cache.NewCacheRedis("tcp", "127.0.0.1:6379", 10)

	// New order controller.
	createOrderCntrlr := rest.NewCreateOrder(actions.NewCreateOrder(repo), respndr)
	retrieveOrderCntrlr := rest.NewRetrieveOrder(actions.NewRetrieveOrder(read.NewOrderFinderById(mysqlDb, read.NewOrderItemFinderById(mysqlDb))), c, respndr)

	// Router and server.
	r := mux.NewRouter()
	r.HandleFunc("/create", createOrderCntrlr.Create)
	r.HandleFunc("/retrieve/{uuid}", retrieveOrderCntrlr.Retrieve)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
