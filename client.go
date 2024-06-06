package main

import (
	"net/http"
)

const (
	endpointBigMessage = "/hello"
)

func registerHandlers(mux *http.ServeMux) {
	mux.Handle(endpointBigMessage, handleHello())
}

func New() http.Handler {
	mux := http.NewServeMux()

	// Register Handlers.
	registerHandlers(mux)

	var handler http.Handler = mux

	// Apply Common Middlewares.

	return handler
}
