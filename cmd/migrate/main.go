package main

import (
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	goMigrateMySql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	driver, err := goMigrateMySql.WithInstance(db, &goMigrateMySql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
