package log

import (
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func SetupLogger(logLevel string) {
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		zlog.Warn().Msg("Failed to parse log level, setting to debug level")
		level = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(level)
	zlog.Info().Msg("Logger setup complete")
}
