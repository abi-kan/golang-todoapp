package tasks_postgres

import (
	"context"
	"errors"
	"fmt"

	core_postgres "github.com/abi-kan/golang-todoapp/internal/core/adapter/postgres"
	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
)

func (p *Pool) CreateTask(
	ctx context.Context,
	task domain.Task,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, p.pool.OpTimeout())
	defer cancel()

	const sql = `
		INSERT INTO todoapp.tasks (title, description, completed, created_at, completed_at, author_user_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id;
		`
	row := p.pool.QueryRow(
		ctx,
		sql,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.CompletedAt,
		task.AuthorUserID,
	)
	newTask, err := scanTaskRow(row)
	if err != nil {
		if errors.Is(err, core_postgres.ErrViolatesForeignKey) {
			return domain.Task{}, fmt.Errorf(
				"%v: user with id='%d': %w",
				err,
				task.AuthorUserID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	return newTask, nil
}
