package users_dto

type GetUsersInput struct {
	Limit  *int
	Offset *int
}

type GetUsersOutput []User
