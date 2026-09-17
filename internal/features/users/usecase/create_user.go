package users_usecase

import (
	"context"
	"fmt"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	users_dto "github.com/abi-kan/golang-todoapp/internal/features/users/dto"
)

func (u *Users) CreateUser(
	ctx context.Context,
	input users_dto.CreateUserInput,
) (users_dto.CreateUserOutput, error) {
	user, err := domain.NewUser(input.FullName, input.PhoneNumber)
	if err != nil {
		return users_dto.CreateUserOutput{}, fmt.Errorf("validate user domain: %w", err)
	}

	newUser, err := u.postgres.CreateUser(ctx, user)
	if err != nil {
		return users_dto.CreateUserOutput{}, fmt.Errorf("create user: %w", err)
	}

	output := users_dto.CreateUserOutput(
		users_dto.NewDTOUserFromDomain(newUser),
	)

	return output, nil
}
