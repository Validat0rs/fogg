package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Validat0rs/fogg/pkg/fogg"
)

var requiredEnv = []string{
	"FOGG_PORT",
	"API_HOST",
	"RPC_HOST",
}

func main() {
	if err := checkEnv(); err != nil {
		log.Fatal(err)
	}

	_fogg := fogg.NewFogg()
	_fogg.SetHandlers()
	_fogg.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	_fogg.Stop()
}

func checkEnv() error {
	for _, envVar := range requiredEnv {
		if os.Getenv(envVar) == "" {
			return fmt.Errorf("%s is not set", envVar)
		}
	}

	return nil
}
