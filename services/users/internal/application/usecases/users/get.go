package users

import (
	"context"

	"github.com/sssoultrix/event-go/services/users/internal/domain"
)

func (uc *usersUseCase) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return uc.repo.GetByEmail(ctx, email)
}

func (uc *usersUseCase) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return uc.repo.GetByID(ctx, id)
}
