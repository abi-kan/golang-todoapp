package users_usecase

import (
	"context"
	"fmt"

	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	"go.uber.org/zap"
)

func (u *Users) DeleteUser(
	ctx context.Context,
	id int,
) error {
	if err := u.postgres.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	if err := u.redis.DeleteUser(ctx, id); err != nil {
		logger := core_logger.FromContext(ctx)
		logger.Warn("cache delete user failed",
			zap.Int("id", id),
			zap.Error(err),
		)
	}

	return nil
}
