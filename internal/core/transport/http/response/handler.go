package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	"go.uber.org/zap"
)

type ResponseHandler struct {
	logger         *core_logger.Logger
	responseWriter http.ResponseWriter
}

func NewResponseHandler(
	logger *core_logger.Logger,
	responseWriter http.ResponseWriter,
) *ResponseHandler {
	return &ResponseHandler{
		logger:         logger,
		responseWriter: responseWriter,
	}
}

func (h *ResponseHandler) JSONResponse(
	responseBody any,
	statusCode int,
) {
	h.responseWriter.WriteHeader(statusCode)
	if err := json.NewEncoder(h.responseWriter).Encode(responseBody); err != nil {
		h.logger.Error("Write HTTP response", zap.Error(err))
	}
}

func (h *ResponseHandler) NoContentResponse() {
	h.responseWriter.WriteHeader(http.StatusNoContent)
}

func (h *ResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.logger.Warn
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.logger.Debug
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.logger.Warn
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.logger.Error
	}

	logFunc(msg, zap.Error(err))
	h.errorResponse(statusCode, err, msg)
}

func (h *ResponseHandler) PanicResponse(anyPanic any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", anyPanic)

	h.logger.Error(msg, zap.Error(err))
	h.errorResponse(statusCode, err, msg)
}

func (h *ResponseHandler) errorResponse(
	statusCode int,
	err error,
	msg string,
) {
	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	h.JSONResponse(response, statusCode)
}
