package users_postgres

import (
	"context"
	"errors"
	"fmt"

	core_postgres "github.com/abi-kan/golang-todoapp/internal/core/adapter/postgres"
	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
)

func (p *Pool) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, p.pool.OpTimeout())
	defer cancel()

	const sql = `
		SELECT id, version, full_name, phone_number
		FROM todoapp.users
		WHERE id=$1;
		`
	row := p.pool.QueryRow(ctx, sql, id)
	user, err := scanUserRow(row)
	if err != nil {
		if errors.Is(err, core_postgres.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return user, nil
}
