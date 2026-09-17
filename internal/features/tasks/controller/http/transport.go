package tasks_http

import (
	"net/http"

	core_http_server "github.com/abi-kan/golang-todoapp/internal/core/transport/http/server"
	tasks_usecase "github.com/abi-kan/golang-todoapp/internal/features/tasks/usecase"
)

type Handler struct {
	service *tasks_usecase.Tasks
}

func NewHandler(service *tasks_usecase.Tasks) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Routes() []core_http_server.Route {
	createTask := core_http_server.NewRoute(
		http.MethodPost,
		"/tasks",
		h.CreateTask,
	)
	getTasks := core_http_server.NewRoute(
		http.MethodGet,
		"/tasks",
		h.GetTasks,
	)
	getTask := core_http_server.NewRoute(
		http.MethodGet,
		"/tasks/{id}",
		h.GetTask,
	)
	deleteTask := core_http_server.NewRoute(
		http.MethodDelete,
		"/tasks/{id}",
		h.DeleteTask,
	)
	patchTask := core_http_server.NewRoute(
		http.MethodPatch,
		"/tasks/{id}",
		h.PatchTask,
	)

	return []core_http_server.Route{
		createTask,
		getTasks,
		getTask,
		deleteTask,
		patchTask,
	}
}
