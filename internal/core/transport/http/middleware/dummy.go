package core_http_middleware

import (
	"fmt"
	"net/http"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
)

func Dummy(s string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.FromContext(ctx)

			logger.Debug(fmt.Sprintf("-> Before: %s", s))

			next.ServeHTTP(w, r)

			logger.Debug(fmt.Sprintf("<- After: %s", s))
		})
	}
}
