package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func Info(msg string, keysAndValues ...interface{}) {
	log.Info().Fields(keysAndValues).Msg(msg)
}

func Debug(msg string, keysAndValues ...interface{}) {
	log.Debug().Fields(keysAndValues).Msg(msg)
}

func Error(msg string, keysAndValues ...interface{}) {
	log.Error().Fields(keysAndValues).Msg(msg)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	log.Fatal().Fields(keysAndValues).Msg(msg)
}
