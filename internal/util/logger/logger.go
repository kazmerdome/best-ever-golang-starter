package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	// levels
	TraceLevel string = "trace"
	DebugLevel string = "debug"
	InfoLevel  string = "info"
	WarnLevel  string = "warn"
	FatalLevel string = "fatal"

	// environments
	Development string = "development"
	Production  string = "production"

	// default level
	DefaultLevelStr       string        = TraceLevel
	DefaultLevelZeroLevel zerolog.Level = zerolog.TraceLevel
)

func InitLogger(level, environment string) {

	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Set default level to trace
	if level == "" {
		level = DefaultLevelStr
	}

	// Parse level string to zerolog.Level
	l, err := zerolog.ParseLevel(level)
	if err != nil {
		log.Error().Err(err).Msg("")
		l = DefaultLevelZeroLevel
	}

	// Setting up level
	zerolog.SetGlobalLevel(l)

	// Init Logger
	var logger zerolog.Logger

	// Add caller if level == debug
	if l == zerolog.DebugLevel {
		logger = log.With().Caller().Logger()
	} else {
		logger = log.With().Logger()
	}

	// Setting up pretty logging for development
	if environment == Development {
		log.Logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		log.Logger = logger
	}
}
