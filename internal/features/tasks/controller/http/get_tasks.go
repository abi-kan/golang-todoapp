package tasks_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_request "github.com/abi-kan/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/abi-kan/golang-todoapp/internal/core/transport/http/response"
	tasks_dto "github.com/abi-kan/golang-todoapp/internal/features/tasks/dto"
)

func (h *Handler) GetTasks(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(logger, rw)

	input, err := getTasksInputFromRequest(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get query params",
		)
		return
	}

	output, err := h.service.GetTasks(ctx, input)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get tasks",
		)
		return
	}

	responseHandler.JSONResponse(output, http.StatusOK)
}

func getTasksInputFromRequest(r *http.Request) (tasks_dto.GetTasksInput, error) {
	const (
		userIDKey = "user_id"
		limitKey  = "limit"
		offsetKey = "offset"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDKey)
	if err != nil {
		return tasks_dto.GetTasksInput{}, fmt.Errorf("get '%s' query param: %w", userIDKey, err)
	}

	limit, err := core_http_request.GetIntQueryParam(r, limitKey)
	if err != nil {
		return tasks_dto.GetTasksInput{}, fmt.Errorf("get '%s' query param: %w", limitKey, err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetKey)
	if err != nil {
		return tasks_dto.GetTasksInput{}, fmt.Errorf("get '%s' query param: %w", offsetKey, err)
	}

	return tasks_dto.GetTasksInput{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	}, nil
}
