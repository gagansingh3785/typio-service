package main

import (
	"github.com/gagansingh3785/typio-service/appcontext"
	"github.com/gagansingh3785/typio-service/config"
	"github.com/gagansingh3785/typio-service/log"
)

func setupTypingService() error {
	// setup configuration
	cfg, err := config.SetupConfig()
	if err != nil {
		return err
	}

	// setup logger
	log.SetupLogger(cfg.Logger.Level)

	// setup application context
	appcontext.Initiate(cfg)

	return nil
}
