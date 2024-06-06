package main

import (
	"net"
	"net/http"
	"time"
)

func main() {
	server := New()
	httpServer := &http.Server{
		Addr:              net.JoinHostPort("localhost", "8080"),
		Handler:           server,
		ReadHeaderTimeout: 10 * time.Second,
	}

	httpServer.ListenAndServe()
}
