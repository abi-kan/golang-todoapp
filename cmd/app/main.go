package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_postgres "github.com/abi-kan/golang-todoapp/internal/core/adapters/postgres"
	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_middleware "github.com/abi-kan/golang-todoapp/internal/core/transport/http/middleware"
	core_http_server "github.com/abi-kan/golang-todoapp/internal/core/transport/http/server"
	users_postgres "github.com/abi-kan/golang-todoapp/internal/features/users/adapter/postgress"
	users_http "github.com/abi-kan/golang-todoapp/internal/features/users/controller/http"
	users_usecase "github.com/abi-kan/golang-todoapp/internal/features/users/usecase"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGALRM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing postgres connection pool")

	pool, err := core_postgres.NewDefaultPool(
		ctx,
		core_postgres.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	httpServer := createServer(pool, logger)
	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}

func createServer(
	pool core_postgres.Pool,
	logger *core_logger.Logger,
) *core_http_server.Server {
	logger.Debug("initializing feature", zap.String("feature", "users"))

	usersUseCase := users_usecase.NewUsers(
		users_postgres.NewPool(pool),
	)
	usersTransportHTTP := users_http.NewHandler(usersUseCase)

	logger.Debug("initializing HTTP server")

	httpServer := core_http_server.NewServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.APIVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	httpServer.RegisterAPIRouters(apiVersionRouter)

	return httpServer
}
