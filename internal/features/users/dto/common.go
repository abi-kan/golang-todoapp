package users_dto

import "github.com/abi-kan/golang-todoapp/internal/core/domain"

type User struct {
	ID          int     `json:"id" example:"10"`
	Version     int     `json:"version" example:"3"`
	FullName    string  `json:"full_name" example:"Ivan Ivanov"`
	PhoneNumber *string `json:"phone_number" example:"+79998887766"`
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
