package domain

import "context"

type UsersService interface {
	Register(ctx context.Context, params CreateUserParams) (*User, error)
	Login(ctx context.Context, params LoginParams) (*User, error)
}
