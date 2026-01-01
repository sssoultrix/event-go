package users

import (
	"context"

	"github.com/sssoultrix/event-go/services/users/internal/domain"
)

func (uc *usersUseCase) Create(ctx context.Context, email, password, passwordConfirm string) (*domain.User, error) {
	user, err := domain.NewUser(email, password, passwordConfirm)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
