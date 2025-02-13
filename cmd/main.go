package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/willy-r/ecom-example/cmd/api"
	"github.com/willy-r/ecom-example/config"
	"github.com/willy-r/ecom-example/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPass,
		Addr:                 config.Envs.DBAddr,
		DBName:               config.Envs.DBUser,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	InitStorage(db)

	server := api.NewApiServer(":8080", db)

	if err := server.Start(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

func InitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("could not ping db: %v", err)
	}

	log.Println("connected to db")
}
