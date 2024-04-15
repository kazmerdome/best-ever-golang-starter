package logger_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/logger"
)

func TestLogger(t *testing.T) {
	assert := assert.New(t)

	// When try to initialize logger with an invalid level, It should throw an error and set the level to default [trace] level
	//
	buf := bytes.Buffer{}
	log.Logger = zerolog.New(&buf).With().Timestamp().Logger()
	logger.InitLogger("invalid level", "development")
	assert.Contains(buf.String(), "Unknown Level String: 'invalid level'")
	assert.Equal(zerolog.GlobalLevel(), zerolog.TraceLevel)

	// When try to initialize logger without specified level, It should return with the default [trace] level
	//
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.InitLogger("", "production")
	assert.Equal(zerolog.GlobalLevel(), zerolog.TraceLevel)
	log.Logger = zerolog.Logger{}

	// When initialize logger with specified level, It should set the level globally
	//
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.InitLogger("debug", "development")
	assert.Equal(zerolog.GlobalLevel(), zerolog.DebugLevel)
}
