package tasks_postgres

import (
	"context"
	"errors"
	"fmt"

	core_postgres "github.com/abi-kan/golang-todoapp/internal/core/adapter/postgres"
	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
)

func (p *Pool) PatchTask(
	ctx context.Context,
	id int,
	task domain.Task,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, p.pool.OpTimeout())
	defer cancel()

	const sql = `
			UPDATE todoapp.tasks
			SET 
				title=$1,
		 		description=$2,
				completed=$3,
				completed_at=$4,
				version=version+1
			WHERE id=$5 AND version=$6
			RETURNING 
				id,
			 	version,
				title,
			 	description,
				completed,
			 	created_at,
				completed_at,
			 	author_user_id;
			`
	row := p.pool.QueryRow(
		ctx,
		sql,
		task.Title,
		task.Description,
		task.Completed,
		task.CompletedAt,
		id,
		task.Version,
	)
	task, err := scanTaskRow(row)
	if err != nil {
		if errors.Is(err, core_postgres.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	return task, nil
}
