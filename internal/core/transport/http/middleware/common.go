package core_http_middleware

import (
	"context"
	"net/http"
	"time"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_response "github.com/abi-kan/golang-todoapp/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	requestIDHeader = "X-Request-ID"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIDHeader, requestID)
			w.Header().Set(requestIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			newLog := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)
			ctx := context.WithValue(r.Context(), "log", newLog)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewResponseHandler(
				logger,
				w,
			)

			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(
						p,
						"During hangling a HTTP request got unexpected panic",
					)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.FromContext(ctx)
			customRW := core_http_response.NewResponseWriter(w)

			startTime := time.Now().UTC()
			logger.Debug(
				">>> Incoming HTTP request",
				zap.String("http_method", r.Method),
				zap.Time("time", startTime),
			)
			next.ServeHTTP(customRW, r)
			logger.Debug(
				">>> Done HTTP request",
				zap.Int("status_code", customRW.GetStatusCodeOrPanic()),
				zap.Duration("latency", time.Since(startTime)),
			)
		})
	}
}
