package tasks_postgres

import (
	"time"

	core_postgres "github.com/abi-kan/golang-todoapp/internal/core/adapter/postgres"
	"github.com/abi-kan/golang-todoapp/internal/core/domain"
)

func scanTaskRow(row core_postgres.Row) (domain.Task, error) {
	var dto = struct {
		ID           int
		Version      int
		Title        string
		Description  *string
		Completed    bool
		CreatedAt    time.Time
		CompletedAt  *time.Time
		AuthorUserID int
	}{}
	err := row.Scan(
		&dto.ID,
		&dto.Version,
		&dto.Title,
		&dto.Description,
		&dto.Completed,
		&dto.CreatedAt,
		&dto.CompletedAt,
		&dto.AuthorUserID,
	)
	if err != nil {
		return domain.Task{}, err
	}

	task := domain.Task{
		ID:           dto.ID,
		Version:      dto.Version,
		Title:        dto.Title,
		Description:  dto.Description,
		Completed:    dto.Completed,
		CreatedAt:    dto.CreatedAt,
		CompletedAt:  dto.CompletedAt,
		AuthorUserID: dto.AuthorUserID,
	}

	return task, nil
}
