package main

import (
	"github.com/gagansingh3785/typio-service/config"
	"github.com/gagansingh3785/typio-service/log"
	zlog "github.com/rs/zerolog/log"
)

func setupTypingService() error {
	// setup configuration
	config, err := config.SetupConfig()
	if err != nil {
		return err
	}

	// setup logger
	log.SetupLogger(config.Logger.Level)

	// from here on we can use the zerolog logger
	zlog.Info().Fields(map[string]interface{}{
		"config": config,
	}).Msg("Printing config")

	return nil
}
