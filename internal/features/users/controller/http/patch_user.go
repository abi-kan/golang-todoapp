package users_http

import (
	"net/http"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_request "github.com/abi-kan/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/abi-kan/golang-todoapp/internal/core/transport/http/response"
	users_dto "github.com/abi-kan/golang-todoapp/internal/features/users/dto"
)

func (h *Handler) PatchUser(
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

	var input users_dto.PatchUserInput
	if err := core_http_request.DecodeAndValidateRequest(r, &input); err != nil {
		responseHandler.ErrorResponse(
			err,
			"Failed to decode and validate HTTP request",
		)
		return
	}

	user, err := h.service.PatchUser(ctx, userID, input)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)
		return
	}

	responseHandler.JSONResponse(user, http.StatusOK)
}
