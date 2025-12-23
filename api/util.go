package api

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
)

// RealIPFromContext returns the real request IP from the huma.Context.
func RealIPFromContext(ctx huma.Context) string {
	r, _ := humachi.Unwrap(ctx)
	return r.RemoteAddr
}
