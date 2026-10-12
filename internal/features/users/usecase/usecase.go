package users_usecase

import (
	"context"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
)

type Postgres interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)
	DeleteUser(
		ctx context.Context,
		id int,
	) error
	PatchUser(
		ctx context.Context,
		id int,
		user domain.User,
	) (domain.User, error)
}

type Redis interface {
	SetUser(ctx context.Context, user domain.User) error
	GetUser(ctx context.Context, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type Users struct {
	postgres Postgres
	redis    Redis
}

func NewUsers(
	postgres Postgres,
	redis Redis,
) *Users {
	return &Users{
		postgres: postgres,
		redis:    redis,
	}
}
