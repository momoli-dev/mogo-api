// Package api provides endpoint registration and routing.
package api

import (
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

// API manages endpoints.
type API struct {
	mux    *chi.Mux
	fluent *fluentRegistrar
}

// Tag sets the tag for further endpoints.
func (a *API) Tag(tag string) {
	a.fluent.lastTag = tag
}

// Middlewares sets the middlewares for further endpoints.
func (a *API) Middlewares(middlewares huma.Middlewares) {
	a.fluent.lastMiddlewares = middlewares
}

// GetHTTPHandler returns the whole registered state as an HTTP handler. This can be used to start the HTTP server.
//
//nolint:ireturn // This library API ensures to always return http.Handler.
func (a *API) GetHTTPHandler() http.Handler {
	return a.mux
}

// NewParams to create a new API.
type NewParams struct {
	Title      string
	Version    string
	Origins    []string
	EnableDocs bool
}

// New creates a new API.
func New(params *NewParams) *API {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(httprate.LimitByIP(250, time.Minute))
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(5 * time.Minute))

	// TODO: http logging

	mux.Use(cors.Handler(
		//nolint:exhaustruct // Not all fields are used.
		cors.Options{
			AllowedOrigins:   params.Origins,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           600,
			Debug:            false,
		}))

	humaConfig := huma.DefaultConfig(params.Title, params.Version)
	if !params.EnableDocs {
		humaConfig.DocsPath = ""
	}
	humaAPI := humachi.New(mux, humaConfig)
	humaAPI.UseMiddleware(MetadataMiddleware)

	registrar := newFluentRegistrar(humaAPI)

	return &API{
		fluent: registrar,
		mux:    mux,
	}
}
