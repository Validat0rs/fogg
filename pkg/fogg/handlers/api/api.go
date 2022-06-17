package api

import (
	"github.com/rs/zerolog"
)

type Api struct {
	logger zerolog.Logger
}

func NewApi(log zerolog.Logger) *Api {
	return &Api{
		logger: log,
	}
}
