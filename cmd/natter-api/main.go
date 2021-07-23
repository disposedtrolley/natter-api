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

	c := m.NewChain()
	c = c.With(m.NewEnsureHTTPMethod(http.MethodPost))
	c = c.With(m.NewEnsureContentType("application/json"))
	c = c.With(m.SetJSONResponseHeader)
	c = c.With(m.SetSecurityResponseHeaders)
	mux.Handle("/spaces", c.Wrap(spaces.CreateHandler(db)))

	server := &http.Server{
		Addr:    ":4567",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
