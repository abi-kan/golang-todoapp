package tasks_http

import (
	"net/http"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_request "github.com/abi-kan/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/abi-kan/golang-todoapp/internal/core/transport/http/response"
	tasks_dto "github.com/abi-kan/golang-todoapp/internal/features/tasks/dto"
)

func (h *Handler) PatchTask(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(logger, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'taskID' path value",
		)
		return
	}

	var input tasks_dto.PatchTaskInput
	if err := core_http_request.DecodeAndValidateRequest(r, &input); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	output, err := h.service.PatchTask(ctx, taskID, input)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch task",
		)
		return
	}

	responseHandler.JSONResponse(output, http.StatusOK)
}
