// Package server provides utilities for managing an HTTP server.
package server

import (
	"net/http"
	"os"
	"time"
)

const (
	readTimeout       = 5 * time.Second
	writeTimeout      = 10 * time.Second
	idleTimeout       = 120 * time.Second
	readHeaderTimeout = readTimeout
)

// Server is the handler for managing an HTTP server.
type Server struct {
	instance *http.Server
}

// Params for creating a new server.
type Params struct {
	Addr    string
	Handler http.Handler
}

// NewServer creates a new server instance.
func NewServer(params *Params) *Server {
	if params.Handler == nil {
		panic("Handler cannot be nil")
	}
	return &Server{
		instance: &http.Server{
			Addr:              params.Addr,
			Handler:           params.Handler,
			ReadTimeout:       readTimeout,
			ReadHeaderTimeout: readHeaderTimeout,
			WriteTimeout:      writeTimeout,
			IdleTimeout:       idleTimeout,
		},
	}
}

// Start the server.
func (s *Server) Start() error {
	if err := s.instance.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// MustStart is same as Start but panics on error.
func (s *Server) MustStart() {
	if err := s.Start(); err != nil {
		panic(err)
	}
}

// Shutdown the server.
func (s *Server) Shutdown() error {
	return s.instance.Close()
}

// MustShutdown is same as Shutdown but panics on error.
func (s *Server) MustShutdown() {
	if err := s.Shutdown(); err != nil {
		panic(err)
	}
}

// StartGracefully starts the server and listens for termination signals to gracefully shut it down. Non-blocking.
func (s *Server) StartGracefully() chan os.Signal {
	c := onTerminate(func(_ os.Signal) {
		s.MustShutdown()
	})

	go func() {
		s.MustStart()
	}()

	return c
}
