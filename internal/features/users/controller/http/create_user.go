package users_http

import (
	"net/http"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_request "github.com/abi-kan/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/abi-kan/golang-todoapp/internal/core/transport/http/response"
	users_dto "github.com/abi-kan/golang-todoapp/internal/features/users/dto"
)

// CreateUser godoc
// @Summary Create user
// @Description Create a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param request body users_dto.CreateUserInput true "CreateUser request body"
// @Success 201 {object} users_dto.CreateUserOutput "User has been successfully created"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users [post]
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
			"failed to decode and validate HTTP request",
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
			"failed to create user",
		)
		return
	}

	responseHandler.JSONResponse(output, http.StatusCreated)
}
