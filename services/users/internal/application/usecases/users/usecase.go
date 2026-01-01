package users

import (
	"context"

	"github.com/sssoultrix/event-go/services/users/internal/application/interfaces"
	"github.com/sssoultrix/event-go/services/users/internal/domain"
)

type UsersUseCase interface {
	Create(ctx context.Context, email, password, passwordConfirm string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, id, email string) (*domain.User, error)
	Delete(ctx context.Context, id string) error
	Login(ctx context.Context, email, password string) (*domain.User, error)
}

type usersUseCase struct {
	repo interfaces.UsersRepository
}

func NewUsersUseCase(repo interfaces.UsersRepository) UsersUseCase {
	return &usersUseCase{
		repo: repo,
	}
}
