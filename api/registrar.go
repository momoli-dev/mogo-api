package api

import (
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

type registrar struct {
	humaAPI huma.API
}

func newRegistrar(humaAPI huma.API) *registrar {
	return &registrar{
		humaAPI: humaAPI,
	}
}

func register[TReq, TRes any](r *registrar, e endpoint, h Handler[TReq, TRes]) {
	opid := genOperationID(e)

	huma.Register(
		r.humaAPI,
		//nolint:exhaustruct // We only use these fields.
		huma.Operation{
			OperationID: opid,
			Method:      e.Method,
			Path:        e.Path,
			Tags:        []string{e.Tag},
			Middlewares: e.Middlewares,
			Summary:     e.Title,
		},
		h,
	)
}

func genOperationID(e endpoint) string {
	if strings.Contains(e.Path, "_") {
		panic("path cannot contain underscores, use dashes instead")
	}

	genPath := strings.ReplaceAll(e.Path, "/", "_")
	genMethod := strings.ToLower(e.Method)
	genID := genMethod + genPath
	return genID
}
