//+build go1.8

// Package expvar extends the expvar package in the stdlib.
package expvar

import (
	"expvar"

	"github.com/goph/stdlib/net/http"
)

// RegisterRoutes register pprof routes in an http.HandlerAcceptor.
func RegisterRoutes(h http.HandlerAcceptor) {
	h.Handle("/debug/vars", expvar.Handler())
}
