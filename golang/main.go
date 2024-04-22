package main

import (
	"go-tavern/paths"
	"go-tavern/router"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://"+paths.MigrationsPath,
		"sqlite3://"+paths.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		sourceErr, databaseErr := m.Close()
		if sourceErr != nil || databaseErr != nil {
			log.Fatalf("source error: %v, database error: %v", sourceErr, databaseErr)
		}
	}()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	handler := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", handler))
}
