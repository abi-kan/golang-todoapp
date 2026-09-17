package tasks_usecase

import (
	"context"
	"fmt"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	tasks_dto "github.com/abi-kan/golang-todoapp/internal/features/tasks/dto"
)

func (t *Tasks) CreateTask(
	ctx context.Context,
	input tasks_dto.CreateTaskInput,
) (tasks_dto.CreateTaskOutput, error) {
	task, err := domain.NewTask(
		input.Title,
		input.Description,
		input.AuthorUserID,
	)
	if err != nil {
		return tasks_dto.CreateTaskOutput{}, fmt.Errorf(
			"validate task domain: %w",
			err,
		)
	}

	newTask, err := t.postgres.CreateTask(ctx, task)
	if err != nil {
		return tasks_dto.CreateTaskOutput{}, fmt.Errorf(
			"create task: %w",
			err,
		)
	}

	output := tasks_dto.CreateTaskOutput(
		tasks_dto.NewDTOTaskFromDomain(newTask),
	)

	return output, nil
}
