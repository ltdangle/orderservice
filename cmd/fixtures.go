package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"orders/model/write"
	"os"
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

	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Truncate tables
	db.Exec("TRUNCATE TABLE orders")
	db.Exec("TRUNCATE TABLE order_items")
	db.Exec("TRUNCATE TABLE payments")

	// Create fixtures.
	orderUuid := "efc144e4"
	db.Create(&write.Order{Uuid: orderUuid, Status: "completed"})
	//db.Create(&write.Payment{Uuid: "8c69", OrderId: orderUuid, PaymentId: "8d7ksk83", Date: time.Now()})
	db.Create(&write.OrderItem{Uuid: "3abb", OrderId: orderUuid, ProductId: "4ef6", Title: "Product one", Description: "Product one description", Price: 1000})
	db.Create(&write.OrderItem{Uuid: "3abc", OrderId: orderUuid, ProductId: "4efc", Title: "Product two", Description: "Product two description", Price: 1000})
	db.Create(&write.OrderItem{Uuid: "3abd", OrderId: orderUuid, ProductId: "4efd", Title: "Product three", Description: "Product three description", Price: 1000})

}
