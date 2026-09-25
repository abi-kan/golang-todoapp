package tasks_postgres

import (
	"context"
	"errors"
	"fmt"

	core_postgres "github.com/abi-kan/golang-todoapp/internal/core/adapter/postgres"
	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
)

func (p *Pool) GetTask(
	ctx context.Context,
	id int,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, p.pool.OpTimeout())
	defer cancel()

	const sql = `
			SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
			FROM todoapp.tasks
			WHERE id=$1;
			`
	row := p.pool.QueryRow(ctx, sql, id)
	task, err := scanTaskRow(row)
	if err != nil {
		if errors.Is(err, core_postgres.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}

		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	return task, nil
}
