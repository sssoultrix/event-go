package users

import (
	"context"
	"fmt"

	"github.com/sssoultrix/event-go/services/users/internal/domain"
)

func (uc *usersUseCase) Update(ctx context.Context, id, email string) (*domain.User, error) {
	user, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	user.Email = email

	if err := uc.repo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}
