package main

import (
	_ "embed"
	"log"

	"github.com/disposedtrolley/natter-api/internal/db"
)

func main() {
	db, err := db.NewInMemory()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
