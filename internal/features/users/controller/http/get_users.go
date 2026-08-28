package users_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_response "github.com/abi-kan/golang-todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/abi-kan/golang-todoapp/internal/core/transport/http/utils"
)

func (h *Handler) GetUsers(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(logger, rw)

	queryParams, err := getUsersQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'queryParams'",
		)
		return
	}

	users, err := h.service.GetUsers(ctx, queryParams.limit, queryParams.offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get users",
		)
		return
	}

	responseHandler.JSONResponse(users, http.StatusOK)
}

type queryParams struct {
	limit  *int
	offset *int
}

func getUsersQueryParams(r *http.Request) (queryParams, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return queryParams{}, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return queryParams{}, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return queryParams{
		limit:  limit,
		offset: offset,
	}, nil
}
