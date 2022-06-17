package types

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type HTTP struct {
	Router *mux.Router
	Server *http.Server
	Client *http.Client
}

type Monitoring struct {
	Logger zerolog.Logger
}

type Fogg struct {
	Secure     bool
	HTTP       HTTP
	Monitoring Monitoring
}
