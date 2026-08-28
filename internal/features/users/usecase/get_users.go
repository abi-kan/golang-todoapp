package users_usecase

import (
	"context"
	"fmt"

	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
	users_dto "github.com/abi-kan/golang-todoapp/internal/features/users/dto"
)

func (u *Users) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) (users_dto.GetUsersOutput, error) {
	if err := validateGetUsersQuery(limit, offset); err != nil {
		return nil, err
	}

	users, err := u.postgres.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users from repository: %w", err)
	}

	output := users_dto.GetUsersOutput(
		users_dto.NewDTOUsersFromDomain(users),
	)

	return output, nil
}

func validateGetUsersQuery(limit, offset *int) error {
	if limit != nil && *limit < 0 {
		return fmt.Errorf(
			"'limit' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return fmt.Errorf(
			"'offset' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
