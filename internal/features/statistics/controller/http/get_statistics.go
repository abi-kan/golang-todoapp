package statistics_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_request "github.com/abi-kan/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/abi-kan/golang-todoapp/internal/core/transport/http/response"
	statistics_dto "github.com/abi-kan/golang-todoapp/internal/features/statistics/dto"
)

func (h *Handler) GetStatistics(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHandler(logger, rw)

	input, err := getStatisticsInputFromRequest(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get query params",
		)
		return
	}

	output, err := h.service.GetStatistics(ctx, input)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get statistics",
		)
		return
	}

	responseHandler.JSONResponse(output, http.StatusOK)
}

func getStatisticsInputFromRequest(r *http.Request) (statistics_dto.GetStatisticsInput, error) {
	const (
		userIDKey = "user_id"
		fromKey   = "from"
		toKey     = "to"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDKey)
	if err != nil {
		return statistics_dto.GetStatisticsInput{}, fmt.Errorf(
			"get '%s' query param: %w", userIDKey, err,
		)
	}

	from, err := core_http_request.GetDateQueryParam(r, fromKey)
	if err != nil {
		return statistics_dto.GetStatisticsInput{}, fmt.Errorf(
			"get '%s' query param: %w", fromKey, err,
		)
	}

	to, err := core_http_request.GetDateQueryParam(r, toKey)
	if err != nil {
		return statistics_dto.GetStatisticsInput{}, fmt.Errorf(
			"get '%s' query param: %w", toKey, err,
		)
	}

	return statistics_dto.GetStatisticsInput{
		UserID: userID,
		From:   from,
		To:     to,
	}, nil
}
