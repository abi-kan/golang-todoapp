package tasks_usecase

import (
	"context"
	"fmt"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	tasks_dto "github.com/abi-kan/golang-todoapp/internal/features/tasks/dto"
)

func (t *Tasks) PatchTask(
	ctx context.Context,
	id int,
	input tasks_dto.PatchTaskInput,
) (tasks_dto.PatchTaskOutput, error) {
	task, err := t.postgres.GetTask(ctx, id)
	if err != nil {
		return tasks_dto.PatchTaskOutput{}, fmt.Errorf("get task: %w", err)
	}

	patch, err := domain.NewTaskPatch(
		input.Title.Nullable,
		input.Description.Nullable,
		input.Completed.Nullable,
	)
	if err != nil {
		return tasks_dto.PatchTaskOutput{}, fmt.Errorf(
			"domain patch create: %w",
			err,
		)
	}

	if err := task.ApplyPatch(patch); err != nil {
		return tasks_dto.PatchTaskOutput{}, fmt.Errorf(
			"apply task patch: %w",
			err,
		)
	}

	patchedTask, err := t.postgres.PatchTask(ctx, id, task)
	if err != nil {
		return tasks_dto.PatchTaskOutput{}, fmt.Errorf(
			"patch task: %w",
			err,
		)
	}

	output := tasks_dto.PatchTaskOutput(tasks_dto.NewDTOTaskFromDomain(patchedTask))

	return output, nil
}
