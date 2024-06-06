package main

import (
	"net"
	"net/http"
	"time"

	"github.com/gnaivex/cdcdtest/server"

	"github.com/gnaivex/tools/log"
)

func main() {
	log.Debugf("starting server")

	srv := server.New()
	httpServer := &http.Server{
		Addr:              net.JoinHostPort("localhost", "8080"),
		Handler:           srv,
		ReadHeaderTimeout: 10 * time.Second,
	}

	httpServer.ListenAndServe()
}
