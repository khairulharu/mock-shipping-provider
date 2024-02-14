package main

import (
	"database/sql"
	"fmt"
	"log"
	"mock-shipping-provider/repository/orderlog"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	database, err := sql.Open("sqlite3", "orderlog.db")
	if err != nil {
		log.Fatalf("error When connenct to sqlite3 get: %v", err.Error())
	}

	defer func() {
		err := database.Close()
		if err != nil {
			log.Fatalf("error; %v", err.Error())
		}
	}()

	if err := database.Ping(); err != nil {
		log.Fatalf("err when ping to databse get: %v", err.Error())
	}

	fmt.Println("running")
	//for test so change with this or delete

	//orderLogRepository := orderlog.New(database)

	_ = orderlog.New(database)
}
