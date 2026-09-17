package users_usecase

import (
	"context"
	"fmt"

	users_dto "github.com/abi-kan/golang-todoapp/internal/features/users/dto"
)

func (u *Users) GetUser(
	ctx context.Context,
	id int,
) (users_dto.GetUserOutput, error) {
	user, err := u.postgres.GetUser(ctx, id)
	if err != nil {
		return users_dto.GetUserOutput{}, fmt.Errorf(
			"get user from repository: %w",
			err,
		)
	}

	output := users_dto.GetUserOutput(users_dto.NewDTOUserFromDomain(user))

	return output, nil
}
