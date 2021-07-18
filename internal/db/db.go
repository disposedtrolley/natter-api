package db

import (
	"database/sql"
	_ "embed"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schema string

// NewInMemory returns a connection to an in-memory SQLite database
// initialised with the schema defined in package db's schema.sql file.
func NewInMemory() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:natter.db?cache=shared&mode=memory")
	if err != nil {
		return nil, fmt.Errorf("open in-memory SQLite connection: %+v", err)
	}

	if err := createTables(db, schema); err != nil {
		return nil, fmt.Errorf("initialise DB schema: %+v", err)
	}

	return db, nil
}

func createTables(db *sql.DB, schema string) error {
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
