package core_pgx

import (
	"errors"

	core_postgres "github.com/abi-kan/golang-todoapp/internal/core/adapter/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRows struct {
	pgx.Rows
}

type pgxRow struct {
	pgx.Row
}

func (r pgxRow) Scan(dest ...any) error {
	err := r.Row.Scan(dest...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return core_postgres.ErrNoRows
		}
	}

	return err
}

type pgxCommandTag struct {
	pgconn.CommandTag
}
