package core_http_server

import (
	"net/http"

	core_http_middleware "github.com/abi-kan/golang-todoapp/internal/core/transport/http/middleware"
)

type Route struct {
	Method      string
	Path        string
	handler     http.HandlerFunc
	middlewares []core_http_middleware.Middleware
}

func NewRoute(
	method string,
	path string,
	handler http.HandlerFunc,
	middlewares ...core_http_middleware.Middleware,
) Route {
	return Route{
		Method:      method,
		Path:        path,
		handler:     handler,
		middlewares: middlewares,
	}
}

func (r *Route) Handler() http.Handler {
	return core_http_middleware.ChainMiddlewares(
		r.handler,
		r.middlewares...,
	)
}
