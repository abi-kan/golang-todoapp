package tasks_usecase

import (
	"context"
	"fmt"

	tasks_dto "github.com/abi-kan/golang-todoapp/internal/features/tasks/dto"
)

func (t *Tasks) GetTask(
	ctx context.Context,
	id int,
) (tasks_dto.GetTaskOutput, error) {
	task, err := t.postgres.GetTask(ctx, id)
	if err != nil {
		return tasks_dto.GetTaskOutput{}, fmt.Errorf(
			"get task from repository: %w",
			err,
		)
	}

	output := tasks_dto.GetTaskOutput(tasks_dto.NewDTOTaskFromDomain(task))

	return output, nil
}
