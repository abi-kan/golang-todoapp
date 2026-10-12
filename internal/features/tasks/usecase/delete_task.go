package tasks_usecase

import (
	"context"
	"fmt"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	"go.uber.org/zap"
)

func (t *Tasks) DeleteTask(
	ctx context.Context,
	id int,
) error {
	if err := t.postgres.DeleteTask(ctx, id); err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	if err := t.redis.DeleteTask(ctx, id); err != nil {
		logger := core_logger.FromContext(ctx)
		logger.Warn("cache delete task failed",
			zap.Int("id", id),
			zap.Error(err),
		)
	}

	return nil
}
