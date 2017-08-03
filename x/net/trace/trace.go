package trace

import (
	"net/http"

	_http "github.com/goph/stdlib/net/http"
	"golang.org/x/net/trace"
)

// RegisterRoutes register pprof routes in an http.HandlerAcceptor.
func RegisterRoutes(h _http.HandlerAcceptor) {
	h.HandleFunc("/debug/requests", Traces)
	h.HandleFunc("/debug/events", Events)
}

// TODO: use Traces and Events functions from trace package as soon as they are available.

// Traces responds with traces from the program.
// The package initialization registers it in http.DefaultServeMux
// at /debug/requests.
//
// It performs authorization by running AuthRequest.
func Traces(w http.ResponseWriter, req *http.Request) {
	any, sensitive := trace.AuthRequest(req)
	if !any {
		http.Error(w, "not allowed", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	trace.Render(w, req, sensitive)
}

// Events responds with a page of events collected by EventLogs.
// The package initialization registers it in http.DefaultServeMux
// at /debug/events.
//
// It performs authorization by running AuthRequest.
func Events(w http.ResponseWriter, req *http.Request) {
	any, sensitive := trace.AuthRequest(req)
	if !any {
		http.Error(w, "not allowed", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	trace.RenderEvents(w, req, sensitive)
}
