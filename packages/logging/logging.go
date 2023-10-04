package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	service string
}

func NewLogger(serviceName string) *Logger {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	return &Logger{
		service: serviceName,
	}
}

func (logger *Logger) Trace(format string, args ...any) {
	log.Trace().Str("service", logger.service).Msgf(format, args...)
}

func (logger *Logger) Fatal(err error) {
	log.Panic().Str("service", logger.service).Err(err).Msg("")
}

func (logger *Logger) Print(format string, args ...any) {
	log.Info().Str("service", logger.service).Msgf(format, args...)
}
