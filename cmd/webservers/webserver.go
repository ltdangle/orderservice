package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("../../.env")
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

	// Mysql orm connection.
	mysqlDb, errr := sql.Open("mysql", mysqlDsn)
	if errr != nil {
		log.Fatal(errr.Error())
	}

	app := buildDependencies(mysqlDb, orm, "2006-01-02 15:04:05", "http://checkout.url")
	app.run()
}
