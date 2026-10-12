package users_usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	users_dto "github.com/abi-kan/golang-todoapp/internal/features/users/dto"
	"go.uber.org/zap"
)

func (u *Users) GetUser(
	ctx context.Context,
	id int,
) (users_dto.GetUserOutput, error) {
	logger := core_logger.FromContext(ctx)

	user, ok := u.getUserFromCache(ctx, logger, id)
	if ok {
		output := users_dto.GetUserOutput(users_dto.NewUserDTOFromDomain(user))
		return output, nil
	}

	user, err := u.getUserFromPostgres(ctx, logger, id)
	if err != nil {
		return users_dto.GetUserOutput{}, err
	}

	output := users_dto.GetUserOutput(users_dto.NewUserDTOFromDomain(user))

	return output, nil
}

func (u *Users) getUserFromCache(
	ctx context.Context,
	logger *core_logger.Logger,
	id int,
) (domain.User, bool) {
	user, err := u.redis.GetUser(ctx, id)
	if err == nil {
		return user, true
	}

	if errors.Is(err, core_errors.ErrNotFound) == false {
		logger.Warn("cache get user failed",
			zap.Int("id", id),
			zap.Error(err),
		)
	}

	return domain.User{}, false
}

func (u *Users) getUserFromPostgres(
	ctx context.Context,
	logger *core_logger.Logger,
	id int,
) (domain.User, error) {
	user, err := u.postgres.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf(
			"get user from repository: %w",
			err,
		)
	}

	if err := u.redis.SetUser(ctx, user); err != nil {
		logger.Warn("cache set user failed",
			zap.Int("id", id),
			zap.Error(err),
		)
	}

	return user, nil
}
