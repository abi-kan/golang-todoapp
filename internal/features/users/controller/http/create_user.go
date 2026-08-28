package users_http

import (
	"net/http"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_request "github.com/abi-kan/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/abi-kan/golang-todoapp/internal/core/transport/http/response"
	users_dto "github.com/abi-kan/golang-todoapp/internal/features/users/dto"
)

func (h *Handler) CreateUser(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(logger, rw)

	var input users_dto.CreateUserInput
	if err := core_http_request.DecodeAndValidateRequest(r, &input); err != nil {
		responseHandler.ErrorResponse(
			err,
			"Failed to decode and validate HTTP request",
		)
		return
	}

	output, err := h.service.CreateUser(
		ctx,
		input,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"Failed to create user",
		)
		return
	}

	responseHandler.JSONResponse(output, http.StatusCreated)
}
