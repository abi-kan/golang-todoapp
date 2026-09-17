package users_usecase

import (
	"context"
	"fmt"

	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
	users_dto "github.com/abi-kan/golang-todoapp/internal/features/users/dto"
)

func (u *Users) GetUsers(
	ctx context.Context,
	input users_dto.GetUsersInput,
) (users_dto.GetUsersOutput, error) {
	if err := validateGetUsersInput(input); err != nil {
		return nil, err
	}

	users, err := u.postgres.GetUsers(ctx, input.Limit, input.Offset)
	if err != nil {
		return nil, fmt.Errorf("get users from repository: %w", err)
	}

	output := users_dto.GetUsersOutput(
		users_dto.NewDTOUsersFromDomain(users),
	)

	return output, nil
}

func validateGetUsersInput(input users_dto.GetUsersInput) error {
	if input.Limit != nil && *input.Limit < 0 {
		return fmt.Errorf(
			"'limit' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if input.Offset != nil && *input.Offset < 0 {
		return fmt.Errorf(
			"'offset' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
