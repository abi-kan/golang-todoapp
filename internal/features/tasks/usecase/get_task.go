package tasks_usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
	core_logger "github.com/abi-kan/golang-todoapp/internal/core/logger"
	tasks_dto "github.com/abi-kan/golang-todoapp/internal/features/tasks/dto"
	"go.uber.org/zap"
)

func (t *Tasks) GetTask(
	ctx context.Context,
	id int,
) (tasks_dto.GetTaskOutput, error) {
	logger := core_logger.FromContext(ctx)

	task, ok := t.getTaskFromCache(ctx, logger, id)
	if ok {
		output := tasks_dto.GetTaskOutput(tasks_dto.NewTaskDTOFromDomain(task))
		return output, nil
	}

	task, err := t.getTaskFromPostgres(ctx, logger, id)
	if err != nil {
		return tasks_dto.GetTaskOutput{}, err
	}

	output := tasks_dto.GetTaskOutput(tasks_dto.NewTaskDTOFromDomain(task))

	return output, nil
}

func (t *Tasks) getTaskFromCache(
	ctx context.Context,
	logger *core_logger.Logger,
	id int,
) (domain.Task, bool) {
	task, err := t.redis.GetTask(ctx, id)
	if err == nil {
		return task, true
	}

	if errors.Is(err, core_errors.ErrNotFound) == false {
		logger.Warn("cache get task failed",
			zap.Int("id", id),
			zap.Error(err),
		)
	}

	return domain.Task{}, false
}

func (t *Tasks) getTaskFromPostgres(
	ctx context.Context,
	logger *core_logger.Logger,
	id int,
) (domain.Task, error) {
	task, err := t.postgres.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf(
			"get task from repository: %w",
			err,
		)
	}

	if err := t.redis.SetTask(ctx, task); err != nil {
		logger.Warn("cache set task failed",
			zap.Int("id", id),
			zap.Error(err),
		)
	}

	return task, nil
}
