package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/abi-kan/golang-todoapp/docs"
	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_middleware "github.com/abi-kan/golang-todoapp/internal/core/transport/http/middleware"
	"github.com/swaggo/http-swagger/v2"
	"go.uber.org/zap"
)

type Server struct {
	mux         *http.ServeMux
	config      Config
	logger      *core_logger.Logger
	middlewares []core_http_middleware.Middleware
}

func NewServer(
	config Config,
	logger *core_logger.Logger,
	middlewares ...core_http_middleware.Middleware,
) *Server {
	return &Server{
		mux:         http.NewServeMux(),
		config:      config,
		logger:      logger,
		middlewares: middlewares,
	}
}

func (s *Server) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)
		s.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router.Handler()),
		)
	}
}

func (s *Server) RegisterSwagger() {
	s.mux.Handle(
		"/swagger/",
		httpSwagger.Handler(
			httpSwagger.URL("/swagger/doc.json"),
			httpSwagger.DefaultModelsExpandDepth(-1),
		),
	)

	s.mux.HandleFunc(
		"/swagger/doc.json",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(docs.SwaggerInfo.ReadDoc()))
		},
	)
}

func (s *Server) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddlewares(
		s.mux,
		s.middlewares...,
	)

	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		s.logger.Warn(
			"Starting HTTP server",
			zap.String("addr", s.config.Addr),
		)

		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) == false {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and server HTTP: %w", err)
		}
	case <-ctx.Done():
		s.logger.Warn("Shutdown HTTP server...")
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			s.config.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		s.logger.Warn("HTTP server stopped")
	}

	return nil
}
