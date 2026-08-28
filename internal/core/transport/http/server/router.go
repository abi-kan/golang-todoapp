package core_http_server

import (
	"fmt"
	"net/http"

	core_http_middleware "github.com/abi-kan/golang-todoapp/internal/core/transport/http/middleware"
)

type APIVersion string

var (
	APIVersion1 = APIVersion("v1")
	APIVersion2 = APIVersion("v2")
	APIVersion3 = APIVersion("v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion  APIVersion
	middlewares []core_http_middleware.Middleware
}

func NewAPIVersionRouter(
	apiVersion APIVersion,
	middlewares ...core_http_middleware.Middleware,
) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:    http.NewServeMux(),
		apiVersion:  apiVersion,
		middlewares: middlewares,
	}
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		r.Handle(pattern, route.Handler())
	}
}

func (r *APIVersionRouter) Handler() http.Handler {
	return core_http_middleware.ChainMiddlewares(
		r,
		r.middlewares...,
	)
}
