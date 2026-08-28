package users_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (p *Pool) PatchUser(
	ctx context.Context,
	id int,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, p.pool.OpTimeout())
	defer cancel()

	const sql = `
		UPDATE todoapp.users
		SET 
			full_name=$1,
	 		phone_number=$2,
			version=version+1
		WHERE id=$3 AND version=$4
		RETURNING id, version, full_name, phone_number;
		`
	row := p.pool.QueryRow(
		ctx,
		sql,
		user.FullName,
		user.PhoneNumber,
		id,
		user.Version,
	)
	user, err := scanUserRow(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return user, nil
}
