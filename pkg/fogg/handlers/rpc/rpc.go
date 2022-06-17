package rpc

import (
	"github.com/rs/zerolog"
)

type Rpc struct {
	logger zerolog.Logger
}

func NewRpc(log zerolog.Logger) *Rpc {
	return &Rpc{
		logger: log,
	}
}
