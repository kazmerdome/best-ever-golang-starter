package server

import (
	"net/http"
)

type Server interface {
	Start() error
	Stop() error
	GetOperations() Operations
}

type Operations interface {
	UseCors()
	UseRecover()
	Get(path string, handler http.HandlerFunc)
	Post(path string, handler http.HandlerFunc)
}
