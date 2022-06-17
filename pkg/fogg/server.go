package fogg

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Validat0rs/fogg/pkg/fogg/types"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type (
	Fogg types.Fogg
)

func NewFogg() *Fogg {
	return &Fogg{
		Secure: false,
		HTTP: types.HTTP{
			Router: mux.NewRouter(),
			Client: &http.Client{Timeout: 5 * time.Second},
		},
		Monitoring: types.Monitoring{
			Logger: log.With().Str("module", "feed").Logger(),
		},
	}
}

func (f *Fogg) Start() {
	f.Monitoring.Logger.Info().Msgf("fogg starting on %v....", ":"+os.Getenv("FOGG_PORT"))
	f.HTTP.Router.Use()
	f.HTTP.Server = &http.Server{
		Addr:         ":" + os.Getenv("FOGG_PORT"),
		Handler:      f.HTTP.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		if err := f.HTTP.Server.ListenAndServe(); err != nil {
			f.Monitoring.Logger.Info().Err(err)
		}
	}()
}

func (f *Fogg) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := f.HTTP.Server.Shutdown(ctx); err != nil {
		f.Monitoring.Logger.Fatal().Err(err).Msg("fogg shutdown")
	}

	f.Monitoring.Logger.Info().Msg("fogg exiting....")
}
