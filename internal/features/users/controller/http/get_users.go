package users_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_request "github.com/abi-kan/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/abi-kan/golang-todoapp/internal/core/transport/http/response"
	users_dto "github.com/abi-kan/golang-todoapp/internal/features/users/dto"
)

// GetUsers godoc
// @Summary Get users
// @Description Get a list of users with optional pagination
// @Tags users
// @Produce json
// @Param limit query int false "Page size"
// @Param offset query int false "Page offset"
// @Success 200 {object} users_dto.GetUsersOutput "The users list has been successfully retrieved"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users [get]
func (h *Handler) GetUsers(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(logger, rw)

	input, err := getUsersInputFromRequest(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get query params",
		)
		return
	}

	output, err := h.service.GetUsers(ctx, input)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get users",
		)
		return
	}

	responseHandler.JSONResponse(output, http.StatusOK)
}

func getUsersInputFromRequest(r *http.Request) (users_dto.GetUsersInput, error) {
	const (
		limitKey  = "limit"
		offsetKey = "offset"
	)

	limit, err := core_http_request.GetIntQueryParam(r, limitKey)
	if err != nil {
		return users_dto.GetUsersInput{}, fmt.Errorf("get '%s' query param: %w", limitKey, err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetKey)
	if err != nil {
		return users_dto.GetUsersInput{}, fmt.Errorf("get '%s' query param: %w", offsetKey, err)
	}

	return users_dto.GetUsersInput{
		Limit:  limit,
		Offset: offset,
	}, nil
}
