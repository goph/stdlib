// Package pprof extends the pprof package in the stdlib.
package pprof

import (
	"net/http/pprof"

	"github.com/goph/stdlib/net/http"
)

// RegisterRoutes register pprof routes in an http.HandlerAcceptor.
func RegisterRoutes(h http.HandlerAcceptor) {
	h.HandleFunc("/debug/pprof/", pprof.Index)
	h.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	h.HandleFunc("/debug/pprof/profile", pprof.Profile)
	h.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	h.HandleFunc("/debug/pprof/trace", pprof.Trace)
}
