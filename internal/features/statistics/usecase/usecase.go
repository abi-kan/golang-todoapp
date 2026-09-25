package statistics_usecase

import (
	"context"
	"time"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
)

type Postgres interface {
	GetTasks(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) ([]domain.Task, error)
}

type Statistics struct {
	postgres Postgres
}

func NewStatistics(postgres Postgres) *Statistics {
	return &Statistics{
		postgres: postgres,
	}
}
