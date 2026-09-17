package tasks_postgres

import (
	"context"
	"fmt"

	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
)

func (p *Pool) DeleteTask(
	ctx context.Context,
	id int,
) error {
	ctx, cancel := context.WithTimeout(ctx, p.pool.OpTimeout())
	defer cancel()

	const sql = `
		DELETE FROM todoapp.tasks
		WHERE id=$1;
		`
	cmdTag, err := p.pool.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"task with id='%d': %w",
			id,
			core_errors.ErrNotFound,
		)
	}

	return nil
}
