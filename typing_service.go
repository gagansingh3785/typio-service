package main

import (
	"github.com/gagansingh3785/typio-service/appcontext"
	"github.com/gagansingh3785/typio-service/config"
	"github.com/gagansingh3785/typio-service/log"
)

func setupTypioService() error {
	// setup configuration
	cfg, err := config.SetupConfig("application")
	if err != nil {
		return err
	}

	// setup logger
	log.SetupLogger(cfg.Logger.Level)

	// setup application context
	if err := appcontext.Initiate(cfg); err != nil {
		return err
	}

	return nil
}
