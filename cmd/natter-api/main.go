package main

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/disposedtrolley/natter-api/internal/db"
	"github.com/disposedtrolley/natter-api/internal/middleware"
	"github.com/disposedtrolley/natter-api/internal/spaces"
)

func main() {
	db, err := db.NewInMemory()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	mux := http.NewServeMux()

	mux.Handle("/spaces", middleware.SetJSONResponseHeader(spaces.CreateHandler(db)))

	server := &http.Server{
		Addr:    ":4567",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
