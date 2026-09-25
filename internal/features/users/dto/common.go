package users_dto

import "github.com/abi-kan/golang-todoapp/internal/core/domain"

type User struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func NewUserDTOFromDomain(user domain.User) User {
	return User{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func NewUsersDTOFromDomain(users []domain.User) []User {
	dtoUsers := make([]User, len(users))
	for index, user := range users {
		dtoUsers[index] = NewUserDTOFromDomain(user)
	}

	return dtoUsers
}
