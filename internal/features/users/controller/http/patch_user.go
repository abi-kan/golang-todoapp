package users_http

import (
	"net/http"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_request "github.com/abi-kan/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/abi-kan/golang-todoapp/internal/core/transport/http/response"
	users_dto "github.com/abi-kan/golang-todoapp/internal/features/users/dto"
)

// PatchUser godoc
// @Summary Change user
// @Description Change a existing user's info in the system.
// @Description ### Logic of the request body fields (Three-state logic):
// @Description 1. **Field is missing**: `phone_number` is ignored, value in the database won't change.
// @Description 2. **Field value is specified**: `"phone_number": "+79998887766"` new value will be saved to the database.
// @Description 3. **Field value is null**: `"phone_number": null` current value in the database will be removed (set to NULL).
// @Description Restrictions: `full_name` can't be set to null
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID of the user to change"
// @Param request body users_dto.PatchUserInput true "PatchUser request body"
// @Success 200 {object} users_dto.PatchUserOutput "User has been successfully changed"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 404 {object} core_http_response.ErrorResponse "User not found"
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /users/{id} [patch]
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
			"failed to decode and validate HTTP request",
		)
		return
	}

	output, err := h.service.PatchUser(ctx, userID, input)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)
		return
	}

	responseHandler.JSONResponse(output, http.StatusOK)
}
