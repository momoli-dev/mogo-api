package api

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type fluentRegistrar struct {
	registrar       *registrar
	lastTag         string
	lastMiddlewares huma.Middlewares
}

func newFluentRegistrar(humaAPI huma.API) *fluentRegistrar {
	return &fluentRegistrar{
		registrar:       newRegistrar(humaAPI),
		lastTag:         "unset",
		lastMiddlewares: huma.Middlewares{},
	}
}

// Get registers a GET endpoint with fluent API.
func Get[TReq, TRes any](api *API, path string, handler Handler[TReq, TRes], title string) {
	register(
		api.fluent.registrar,
		endpoint{
			Method:      http.MethodGet,
			Path:        path,
			Title:       title,
			Tag:         api.fluent.lastTag,
			Middlewares: api.fluent.lastMiddlewares,
		},
		handler,
	)
}

// Post registers a POST endpoint with fluent API.
func Post[TReq, TRes any](api *API, path string, handler Handler[TReq, TRes], title string) {
	register(
		api.fluent.registrar,
		endpoint{
			Method:      http.MethodPost,
			Path:        path,
			Title:       title,
			Tag:         api.fluent.lastTag,
			Middlewares: api.fluent.lastMiddlewares,
		},
		handler,
	)
}

// Put registers a PUT endpoint with fluent API.
func Put[TReq, TRes any](api *API, path string, handler Handler[TReq, TRes], title string) {
	register(
		api.fluent.registrar,
		endpoint{
			Method:      http.MethodPut,
			Path:        path,
			Title:       title,
			Tag:         api.fluent.lastTag,
			Middlewares: api.fluent.lastMiddlewares,
		},
		handler,
	)
}

// Delete registers a DELETE endpoint with fluent API.
func Delete[TReq, TRes any](api *API, path string, handler Handler[TReq, TRes], title string) {
	register(
		api.fluent.registrar,
		endpoint{
			Method:      http.MethodDelete,
			Path:        path,
			Title:       title,
			Tag:         api.fluent.lastTag,
			Middlewares: api.fluent.lastMiddlewares,
		},
		handler,
	)
}
