package domain

import (
	"fmt"
	"time"

	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
)

type Task struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}

func NewTask(
	title string,
	description *string,
	authorUserID int,
) (Task, error) {
	task := Task{
		Title:        title,
		Description:  description,
		CreatedAt:    time.Now(),
		AuthorUserID: authorUserID,
	}

	if err := task.validate(); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (t *Task) ApplyPatch(patch TaskPatch) error {
	tmpTask := *t

	if patch.Title.Set {
		tmpTask.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		tmpTask.Description = patch.Description.Value
	}

	if patch.Completed.Set {
		tmpTask.Completed = *patch.Completed.Value
		if tmpTask.Completed {
			completedAt := time.Now()
			tmpTask.CompletedAt = &completedAt
		} else {
			tmpTask.CompletedAt = nil
		}
	}

	if err := tmpTask.validate(); err != nil {
		return fmt.Errorf("validate patched task: %w", err)
	}

	*t = tmpTask

	return nil
}

func (t *Task) validate() error {
	if err := ValidateTaskTitle(t.Title); err != nil {
		return err
	}

	if err := ValidateTaskDescription(t.Description); err != nil {
		return err
	}

	if err := validateTaskCompletion(t.Completed, t.CreatedAt, t.CompletedAt); err != nil {
		return err
	}

	return nil
}

type TaskPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func NewTaskPatch(
	title Nullable[string],
	description Nullable[string],
	completed Nullable[bool],
) (TaskPatch, error) {
	if title.Set && title.Value == nil {
		return TaskPatch{}, fmt.Errorf(
			"'Title' can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if completed.Set && completed.Value == nil {
		return TaskPatch{}, fmt.Errorf(
			"'Completed' can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	patch := TaskPatch{
		Title:       title,
		Description: description,
		Completed:   completed,
	}

	return patch, nil
}
