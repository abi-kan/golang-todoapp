package tasks_postgres

import (
	"context"
	"fmt"

	core_postgres "github.com/abi-kan/golang-todoapp/internal/core/adapter/postgres"
	"github.com/abi-kan/golang-todoapp/internal/core/domain"
)

func (p *Pool) GetTasks(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, p.pool.OpTimeout())
	defer cancel()

	sql := `
			SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
			FROM todoapp.tasks
			%s
			ORDER BY id ASC 
			LIMIT $1
			OFFSET $2;
			`
	args := []any{limit, offset}
	if userID != nil {
		sql = fmt.Sprintf(sql, "WHERE author_user_id=$3")
		args = append(args, userID)
	} else {
		sql = fmt.Sprintf(sql, "")
	}

	rows, err := p.pool.Query(
		ctx,
		sql,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("select tasks: %w", err)
	}

	defer rows.Close()

	return scanGetTasksRows(rows)
}

func scanGetTasksRows(rows core_postgres.Rows) ([]domain.Task, error) {
	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(
			&task.ID,
			&task.Version,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
			&task.AuthorUserID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	return tasks, nil
}
