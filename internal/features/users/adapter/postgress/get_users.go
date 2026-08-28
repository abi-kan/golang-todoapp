package users_postgres

import (
	"context"
	"fmt"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

func (p *Pool) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, p.pool.OpTimeout())
	defer cancel()

	const sql = `
		SELECT id, version, full_name, phone_number
		FROM todoapp.users
		ORDER BY id ASC
		LIMIT $1
		OFFSET $2
		`
	rows, err := p.pool.Query(
		ctx,
		sql,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("select users: %w", err)
	}

	defer rows.Close()

	return scanGetUsersRows(rows)
}

func scanGetUsersRows(rows pgx.Rows) ([]domain.User, error) {
	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.ID,
			&user.Version,
			&user.FullName,
			&user.PhoneNumber,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	return users, nil
}
