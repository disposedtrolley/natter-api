package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Chain struct {
	before []Middleware
}

func NewChain(mw ...Middleware) Chain {
	return Chain{
		before: mw,
	}
}

func (c Chain) Wrap(h http.Handler) http.Handler {
	for _, mw := range c.before {
		h = mw(h)
	}

	return h

}
