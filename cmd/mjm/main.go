package main

import (
	"log"
	"net/http"
	"time"

	"mjm/app"

	"github.com/gobuffalo/buffalo/servers"
)

const (
	// responseTimeout used for request Read, Write and Idle.
	responseTimeout = 30 * time.Second
)

// We initialize some server settings to avoid long running
// requests that would kill our DB.
var server = &servers.Simple{
	Server: &http.Server{
		ReadTimeout:  responseTimeout,
		WriteTimeout: responseTimeout,
		IdleTimeout:  responseTimeout,
	},
}

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Serve(server); err != nil {
		log.Fatal(err)
	}
}
