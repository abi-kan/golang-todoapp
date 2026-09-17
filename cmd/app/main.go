package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_postgres "github.com/abi-kan/golang-todoapp/internal/core/adapter/postgres"
	core_pgx "github.com/abi-kan/golang-todoapp/internal/core/adapter/postgres/pgx"
	core_config "github.com/abi-kan/golang-todoapp/internal/core/config"
	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	core_http_middleware "github.com/abi-kan/golang-todoapp/internal/core/transport/http/middleware"
	core_http_server "github.com/abi-kan/golang-todoapp/internal/core/transport/http/server"
	tasks_postgres "github.com/abi-kan/golang-todoapp/internal/features/tasks/adapter/postgress"
	tasks_http "github.com/abi-kan/golang-todoapp/internal/features/tasks/controller/http"
	tasks_usecase "github.com/abi-kan/golang-todoapp/internal/features/tasks/usecase"
	users_postgres "github.com/abi-kan/golang-todoapp/internal/features/users/adapter/postgress"
	users_http "github.com/abi-kan/golang-todoapp/internal/features/users/controller/http"
	users_usecase "github.com/abi-kan/golang-todoapp/internal/features/users/usecase"
	"go.uber.org/zap"
)

func main() {
	config := core_config.NewConfigMust()
	time.Local = config.TimeZone

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

	logger.Debug("application time zone", zap.Any("zone", time.Local))

	logger.Debug("initializing postgres connection pool")
	pool, err := core_pgx.NewPool(
		ctx,
		core_pgx.NewConfigMust(),
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
	usersTransportHTTP := users_http.NewHandler(
		users_usecase.NewUsers(
			users_postgres.NewPool(pool),
		),
	)

	logger.Debug("initializing feature", zap.String("feature", "tasks"))
	tasksTransportHTTP := tasks_http.NewHandler(
		tasks_usecase.NewTasks(
			tasks_postgres.NewPool(pool),
		),
	)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.APIVersion1)
	apiVersionRouterV1.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRoutes(tasksTransportHTTP.Routes()...)
	httpServer.RegisterAPIRouters(apiVersionRouterV1)

	/*
		// Example of middleware usage on the whole Router.
		apiVersionRouterV2 := core_http_server.NewAPIVersionRouter(
			core_http_server.APIVersion2,
			core_http_middleware.Dummy("api v2 middleware"),
		)
		apiVersionRouterV2.RegisterRoutes(usersTransportHTTP.Routes()...)
		httpServer.RegisterAPIRouters(apiVersionRouterV2)
	*/

	return httpServer
}
