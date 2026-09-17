package users_usecase

import (
	"context"
	"fmt"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	users_dto "github.com/abi-kan/golang-todoapp/internal/features/users/dto"
)

func (u *Users) PatchUser(
	ctx context.Context,
	id int,
	input users_dto.PatchUserInput,
) (users_dto.PatchUserOutput, error) {
	user, err := u.postgres.GetUser(ctx, id)
	if err != nil {
		return users_dto.PatchUserOutput{}, fmt.Errorf("get user: %w", err)
	}

	patch, err := domain.NewUserPatch(
		input.FullName.Nullable,
		input.PhoneNumber.Nullable,
	)
	if err != nil {
		return users_dto.PatchUserOutput{}, fmt.Errorf(
			"domain patch create: %w",
			err,
		)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return users_dto.PatchUserOutput{}, fmt.Errorf(
			"apply user patch: %w",
			err,
		)
	}

	patchedUser, err := u.postgres.PatchUser(ctx, id, user)
	if err != nil {
		return users_dto.PatchUserOutput{}, fmt.Errorf(
			"patch user: %w",
			err,
		)
	}

	output := users_dto.PatchUserOutput(users_dto.NewDTOUserFromDomain(patchedUser))

	return output, nil
}
