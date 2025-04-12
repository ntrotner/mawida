package logger

import (
	"os"

	"template_backend/core/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if config.GlobalConfig.Environment == config.Development {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}
