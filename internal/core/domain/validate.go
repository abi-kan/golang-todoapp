package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
)

func ValidateUserFullName(fullName string) error {
	fullNameLength := len([]rune(fullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf(
			"invalid 'FullName' len: %d: %w",
			fullNameLength,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func ValidateUserPhoneNumber(phoneNumber *string) error {
	if phoneNumber != nil {
		phoneNumberLength := len([]rune(*phoneNumber))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf(
				"invalid 'PhoneNumber' len: %d: %w",
				phoneNumberLength,
				core_errors.ErrInvalidArgument,
			)
		}

		reg := regexp.MustCompile(`^\+[0-9]+$`)
		if reg.MatchString(*phoneNumber) == false {
			return fmt.Errorf(
				"invalid 'PhoneNumber' format: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}
