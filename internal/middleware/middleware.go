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

func (c Chain) With(mw Middleware) Chain {
	c.before = append(c.before, mw)
	return c
}

func (c Chain) Wrap(h http.Handler) http.Handler {
	// Apply middleware in reverse order (so it's executed in
	// the correct order).
	for i := range c.before {
		h = c.before[len(c.before)-1-i](h)
	}

	return h

}
