package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type echoServer struct {
	port   int
	e      *echo.Echo
	logger zerolog.Logger
}

func NewEchoServer(port int) *echoServer {
	return &echoServer{
		port: port,
		e:    echo.New(),
		logger: log.
			With().
			Str("actor", "server/echo").
			Logger(),
	}
}

func (r *echoServer) GetOperations() Operations {
	return newEchoOperations(r.e)
}

func (r *echoServer) Start() error {
	r.logger.
		Info().
		Str("event", "server is starting...").
		Send()
	err := r.e.Start(fmt.Sprintf(":%d", r.port))
	if err != nil && !strings.Contains(err.Error(), "http: Server closed") {
		r.logger.
			Info().
			Err(err).
			Str("event", "starting the server is failed").
			Send()
		return err
	}
	return nil
}

func (r *echoServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.logger.
		Info().
		Str("event", "shutting down the server gracefully...").
		Send()
	err := r.e.Shutdown(ctx)
	if err != nil {
		r.logger.
			Fatal().
			Err(err).
			Str("event", "shutting down the server gracefully failed").
			Send()
		return err
	}
	r.logger.
		Info().
		Str("event", "the server is shut down gracefully").
		Send()
	return nil
}

type echoOperations struct {
	e *echo.Echo
}

func newEchoOperations(e *echo.Echo) *echoOperations {
	return &echoOperations{
		e: e,
	}
}

func (r *echoOperations) UseCors() {
	r.e.Use(middleware.CORS())
}

func (r *echoOperations) UseRecover() {
	r.e.Use(middleware.Recover())
}

func (r *echoOperations) Get(path string, handler http.HandlerFunc) {
	r.e.GET(path, echo.WrapHandler(handler))
}

func (r *echoOperations) Post(path string, handler http.HandlerFunc) {
	r.e.POST(path, echo.WrapHandler(handler))
}
