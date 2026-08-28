package domain

import (
	"fmt"

	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
)

type User struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func NewUser(
	fullName string,
	phoneNumber *string,
) (User, error) {
	user := User{
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}

	if err := user.validate(); err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	tmpUser := *u

	if patch.FullName.Set {
		tmpUser.FullName = *patch.FullName.Value
	}

	if patch.PhoneNumber.Set {
		tmpUser.PhoneNumber = patch.PhoneNumber.Value
	}

	if err := tmpUser.validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmpUser

	return nil
}

func (u *User) validate() error {
	if err := ValidateUserFullName(u.FullName); err != nil {
		return err
	}

	if err := ValidateUserPhoneNumber(u.PhoneNumber); err != nil {
		return err
	}

	return nil
}

type UserPatch struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func NewUserPatch(
	fullName Nullable[string],
	phoneNumber Nullable[string],
) (UserPatch, error) {
	if fullName.Set && fullName.Value == nil {
		return UserPatch{}, fmt.Errorf(
			"'FullName' can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	patch := UserPatch{
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}

	return patch, nil
}
