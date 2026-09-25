package users_postgres

import (
	"context"
	"fmt"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
)

func (p *Pool) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, p.pool.OpTimeout())
	defer cancel()

	const sql = `
		INSERT INTO todoapp.users (full_name, phone_number)
		VALUES ($1, $2)
		RETURNING id, version, full_name, phone_number;
		`
	row := p.pool.QueryRow(
		ctx,
		sql,
		user.FullName,
		user.PhoneNumber,
	)
	newUser, err := scanUserRow(row)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return newUser, nil
}
