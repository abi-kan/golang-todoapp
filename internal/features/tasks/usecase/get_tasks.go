package tasks_usecase

import (
	"context"
	"fmt"

	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
	tasks_dto "github.com/abi-kan/golang-todoapp/internal/features/tasks/dto"
)

func (t *Tasks) GetTasks(
	ctx context.Context,
	input tasks_dto.GetTasksInput,
) (tasks_dto.GetTasksOutput, error) {
	if err := validateGetTasksInput(input); err != nil {
		return nil, err
	}

	tasks, err := t.postgres.GetTasks(ctx, input.UserID, input.Limit, input.Offset)
	if err != nil {
		return nil, fmt.Errorf("get tasks from repository: %w", err)
	}

	output := tasks_dto.GetTasksOutput(
		tasks_dto.NewDTOTasksFromDomain(tasks),
	)

	return output, nil
}

func validateGetTasksInput(input tasks_dto.GetTasksInput) error {
	if input.Limit != nil && *input.Limit < 0 {
		return fmt.Errorf(
			"'limit' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if input.Offset != nil && *input.Offset < 0 {
		return fmt.Errorf(
			"'offset' must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
