package main

import (
	"database/sql"
	_ "embed"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed resources/schema.sql
var schema string

func main() {
	db, err := sql.Open("sqlite3", "file:natter.db?cache=shared&mode=memory")
	if err != nil {
		log.Fatalf("open in-memory SQLite connection: %+v", err)
	}
	defer db.Close()

	if err := createTables(db); err != nil {
		log.Fatalf("initialise DB schema: %+v", err)
	}
}

func createTables(db *sql.DB) error {
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
