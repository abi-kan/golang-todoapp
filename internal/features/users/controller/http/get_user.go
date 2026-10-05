package users_http

import (
	"net/http"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_request "github.com/abi-kan/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/abi-kan/golang-todoapp/internal/core/transport/http/response"
	_ "github.com/abi-kan/golang-todoapp/internal/features/users/dto" // for swagger
)

// GetUser godoc
// @Summary Get user
// @Description Get a concrete user by ID
// @Tags users
// @Produce json
// @Param id path int true "ID of the user to get"
// @Success 200 {object} users_dto.GetUserOutput "User has been successfully found"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "User not found"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users/{id} [get]
func (h *Handler) GetUser(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(logger, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'userID' path value",
		)
		return
	}

	output, err := h.service.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)
		return
	}

	responseHandler.JSONResponse(output, http.StatusOK)
}
