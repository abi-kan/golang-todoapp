package tasks_usecase

import (
	"context"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
)

type Postgres interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)
	GetTasks(
		ctx context.Context,
		userID *int,
		limit *int,
		offset *int,
	) ([]domain.Task, error)
	GetTask(
		ctx context.Context,
		id int,
	) (domain.Task, error)
	DeleteTask(
		ctx context.Context,
		id int,
	) error
	PatchTask(
		ctx context.Context,
		id int,
		task domain.Task,
	) (domain.Task, error)
}

type Tasks struct {
	postgres Postgres
}

func NewTasks(postgres Postgres) *Tasks {
	return &Tasks{
		postgres: postgres,
	}
}
