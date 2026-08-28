package users_dto

import (
	"fmt"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	core_http_types "github.com/abi-kan/golang-todoapp/internal/core/transport/http/types"
)

type PatchUserInput struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

func (input *PatchUserInput) Validate() error {
	if input.FullName.Set {
		if input.FullName.Value == nil {
			return fmt.Errorf("'FullName' can't be null")
		}

		if err := domain.ValidateUserFullName(*input.FullName.Value); err != nil {
			return err
		}
	}

	if input.PhoneNumber.Set {
		if input.PhoneNumber.Value != nil {
			if err := domain.ValidateUserPhoneNumber(input.PhoneNumber.Value); err != nil {
				return err
			}
		}
	}

	return nil
}

type PatchUserOutput User
