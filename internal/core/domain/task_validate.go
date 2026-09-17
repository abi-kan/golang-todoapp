package domain

import (
	"fmt"
	"time"

	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
)

func ValidateTaskTitle(title string) error {
	titleLength := len([]rune(title))
	if titleLength < 1 || titleLength > 100 {
		return fmt.Errorf(
			"invalid 'Title' len: %d: %w",
			titleLength,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func ValidateTaskDescription(description *string) error {
	if description != nil {
		descriptionLength := len([]rune(*description))
		if descriptionLength < 1 || descriptionLength > 1000 {
			return fmt.Errorf(
				"invalid 'Description' len: %d: %w",
				descriptionLength,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}

func validateTaskCompletion(
	completed bool,
	createdAt time.Time,
	completedAt *time.Time,
) error {
	if completed {
		if completedAt == nil {
			return fmt.Errorf(
				"'CompletedAt' can't be nil if 'Completed'=='true': %w",
				core_errors.ErrInvalidArgument,
			)
		}

		if completedAt.Before(createdAt) {
			return fmt.Errorf(
				"'CompletedAt' can't be before 'CreatedAt': %w",
				core_errors.ErrInvalidArgument,
			)
		}
	} else {
		if completedAt != nil {
			return fmt.Errorf(
				"'CompletedAt' must be 'nil' if 'Completed'=='false': %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}
