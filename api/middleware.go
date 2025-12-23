package api

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
)

// CtxMetadataAddr is the context key for the client remote address.
type CtxMetadataAddr struct{}

// CtxMetadataAgent is the context key for the client user agent.
type CtxMetadataAgent struct{}

// MetadataMiddleware adds the client remote address and user agent to the context.
func MetadataMiddleware(ctx huma.Context, next func(huma.Context)) {
	r, _ := humachi.Unwrap(ctx)
	addr := r.RemoteAddr
	agent := r.UserAgent()

	ctx = huma.WithValue(ctx, CtxMetadataAddr{}, addr)
	ctx = huma.WithValue(ctx, CtxMetadataAgent{}, agent)

	next(ctx)
}

// MetadataAddrFromCtx returns the client remote address from the context.
func MetadataAddrFromCtx(ctx context.Context) string {
	addr := ctx.Value(CtxMetadataAddr{})
	if addr == nil {
		return ""
	}

	val, ok := addr.(string)
	if !ok {
		return ""
	}

	return val
}

// MetadataAgentFromCtx returns the client user agent from the context.
func MetadataAgentFromCtx(ctx context.Context) string {
	agent := ctx.Value(CtxMetadataAgent{})
	if agent == nil {
		return ""
	}

	val, ok := agent.(string)
	if !ok {
		return ""
	}

	return val
}
