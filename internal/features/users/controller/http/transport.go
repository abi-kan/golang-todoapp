package users_http

import (
	"net/http"

	core_http_server "github.com/abi-kan/golang-todoapp/internal/core/transport/http/server"
	users_usecase "github.com/abi-kan/golang-todoapp/internal/features/users/usecase"
)

type Handler struct {
	service *users_usecase.Users
}

func NewHandler(service *users_usecase.Users) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Routes() []core_http_server.Route {
	createUser := core_http_server.NewRoute(
		http.MethodPost,
		"/users",
		h.CreateUser,
	)
	getUsers := core_http_server.NewRoute(
		http.MethodGet,
		"/users",
		h.GetUsers,
		// Example of middleware usage on one Route.
		// core_http_middleware.Dummy("get users middleware"),
	)
	getUser := core_http_server.NewRoute(
		http.MethodGet,
		"/users/{id}",
		h.GetUser,
	)
	deleteUser := core_http_server.NewRoute(
		http.MethodDelete,
		"/users/{id}",
		h.DeleteUser,
	)
	patchUser := core_http_server.NewRoute(
		http.MethodPatch,
		"/users/{id}",
		h.PatchUser,
	)

	return []core_http_server.Route{
		createUser,
		getUsers,
		getUser,
		deleteUser,
		patchUser,
	}
}
