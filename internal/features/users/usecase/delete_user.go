package users_usecase

import (
	"context"
	"fmt"
)

func (u *Users) DeleteUser(
	ctx context.Context,
	id int,
) error {
	if err := u.postgres.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}
