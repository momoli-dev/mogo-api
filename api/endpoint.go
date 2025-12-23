package api

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
)

// Handler is a function that handles an endpoint accepting a request and returning a response.
type Handler[TReq, TRes any] func(context.Context, *TReq) (*TRes, error)

// NotImplemented is a generic placeholder handler that fails with a 501 Not Implemented error.
func NotImplemented(_ context.Context, _ *struct{}) (*struct{}, error) {
	return nil, huma.Error501NotImplemented("Not implemented")
}

type endpoint struct {
	Method      string
	Path        string
	Title       string
	Tag         string
	Middlewares huma.Middlewares
}
