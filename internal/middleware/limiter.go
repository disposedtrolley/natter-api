package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/sethvargo/go-limiter"
	"github.com/sethvargo/go-limiter/memorystore"
)

type Limiter struct {
	store limiter.Store
}

func NewLimiter(requests int, interval time.Duration) *Limiter {
	store, err := memorystore.New(&memorystore.Config{
		Tokens:   uint64(requests),
		Interval: interval,
	})

	if err != nil {
		log.Fatalf("create limiter middleware: %+v", err)
	}

	return &Limiter{
		store: store,
	}
}

func (l *Limiter) take(key string) (remaining uint64, limited bool) {
	_, remaining, _, ok, _ := l.store.Take(context.TODO(), key)

	return remaining, ok
}

func (l *Limiter) Handler() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			if _, ok := l.take(ip); !ok {
				http.Error(w, "", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
