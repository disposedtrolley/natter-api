package main

import (
	_ "embed"
	"log"
	"net/http"
	"time"

	"github.com/disposedtrolley/natter-api/internal/db"
	m "github.com/disposedtrolley/natter-api/internal/middleware"
	"github.com/disposedtrolley/natter-api/internal/spaces"
	"github.com/gorilla/mux"
)

func main() {
	db, err := db.NewInMemory()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	limiter := m.NewLimiter(2, time.Second)

	c := m.NewChain()
	c = c.With(m.NewEnsureContentType("application/json"))
	c = c.With(m.SetJSONResponseHeader)
	c = c.With(m.SetSecurityResponseHeaders)

	r := mux.NewRouter()
	r.Use(limiter.Handler())
	r.Handle("/spaces", c.Wrap(spaces.CreateHandler(db))).Methods(http.MethodPost)

	server := &http.Server{
		Addr:    ":4567",
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
