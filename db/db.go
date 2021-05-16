package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Client               *sql.DB
	mysql_users_username = ""
	mysql_users_password = ""
	mysql_users_host     = ""
	mysql_users_schema   = ""
)

func init() {

	// * try using ORM too!

	err := godotenv.Load()

	if err != nil {
		log.Printf("Error loading .env file %s", err.Error())
	}
	log.Print("read .env file successfully")
	mysql_users_username = os.Getenv("mysql_users_username")
	mysql_users_host = os.Getenv("mysql_users_host")
	mysql_users_password = os.Getenv("mysql_users_password")
	mysql_users_schema = os.Getenv("mysql_users_schema")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		mysql_users_username,
		mysql_users_password,
		mysql_users_host,
		mysql_users_schema,
	)

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		// TODO : Handle ERROR
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured!")

}
