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

	// Migrate the schema
	db.AutoMigrate(&write.Order{})
	db.AutoMigrate(&write.OrderItem{})
	db.AutoMigrate(&write.Payment{})

}
