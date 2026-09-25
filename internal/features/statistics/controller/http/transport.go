package statistics_http

import (
	"net/http"

	core_http_server "github.com/abi-kan/golang-todoapp/internal/core/transport/http/server"
	statistics_usecase "github.com/abi-kan/golang-todoapp/internal/features/statistics/usecase"
)

type Handler struct {
	service *statistics_usecase.Statistics
}

func NewHandler(service *statistics_usecase.Statistics) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Routes() []core_http_server.Route {
	getStatistics := core_http_server.NewRoute(
		http.MethodGet,
		"/statistics",
		h.GetStatistics,
	)

	return []core_http_server.Route{
		getStatistics,
	}
}
