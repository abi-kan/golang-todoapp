package users_http

import (
	"net/http"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_response "github.com/abi-kan/golang-todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/abi-kan/golang-todoapp/internal/core/transport/http/utils"
)

func (h *Handler) GetUser(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(logger, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'userID' path value",
		)
		return
	}

	user, err := h.service.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)
		return
	}

	responseHandler.JSONResponse(user, http.StatusOK)
}
