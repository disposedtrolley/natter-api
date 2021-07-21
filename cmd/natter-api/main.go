package main

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/disposedtrolley/natter-api/internal/db"
	m "github.com/disposedtrolley/natter-api/internal/middleware"
	"github.com/disposedtrolley/natter-api/internal/spaces"
)

func main() {
	db, err := db.NewInMemory()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	mux := http.NewServeMux()

	onlyPost := m.NewEnsureHTTPMethod(http.MethodPost)

	mux.Handle("/spaces",
		onlyPost(
			m.SetJSONResponseHeader(
				spaces.CreateHandler(db))))

	server := &http.Server{
		Addr:    ":4567",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
