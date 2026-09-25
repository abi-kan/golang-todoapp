package users_postgres

import (
	core_postgres "github.com/abi-kan/golang-todoapp/internal/core/adapter/postgres"
	"github.com/abi-kan/golang-todoapp/internal/core/domain"
)

func scanUserRow(row core_postgres.Row) (domain.User, error) {
	var dto = struct {
		ID          int
		Version     int
		FullName    string
		PhoneNumber *string
	}{}
	err := row.Scan(
		&dto.ID,
		&dto.Version,
		&dto.FullName,
		&dto.PhoneNumber,
	)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		ID:          dto.ID,
		Version:     dto.Version,
		FullName:    dto.FullName,
		PhoneNumber: dto.PhoneNumber,
	}

	return user, nil
}
