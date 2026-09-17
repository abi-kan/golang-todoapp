package tasks_dto

import (
	"fmt"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	core_http_types "github.com/abi-kan/golang-todoapp/internal/core/transport/http/types"
)

type PatchTaskInput struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed   core_http_types.Nullable[bool]   `json:"completed"`
}

func (input *PatchTaskInput) Validate() error {
	if input.Title.Set {
		if input.Title.Value == nil {
			return fmt.Errorf("'Title' can't be null")
		}

		if err := domain.ValidateTaskTitle(*input.Title.Value); err != nil {
			return err
		}
	}

	if input.Description.Set {
		if input.Description.Value != nil {
			if err := domain.ValidateTaskDescription(input.Description.Value); err != nil {
				return err
			}
		}
	}

	if input.Completed.Set {
		if input.Completed.Value == nil {
			return fmt.Errorf("'Completed' can't be null")
		}
	}

	return nil
}

type PatchTaskOutput Task
