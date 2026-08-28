package users_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_request "github.com/abi-kan/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/abi-kan/golang-todoapp/internal/core/transport/http/response"
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
	const (
		limitKey  = "limit"
		offsetKey = "offset"
	)

	limit, err := core_http_request.GetIntQueryParam(r, limitKey)
	if err != nil {
		return queryParams{}, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetKey)
	if err != nil {
		return queryParams{}, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return queryParams{
		limit:  limit,
		offset: offset,
	}, nil
}
