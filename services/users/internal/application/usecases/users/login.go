package users

import (
	"context"
	"fmt"

	"github.com/sssoultrix/event-go/services/users/internal/domain"
)

func (uc *usersUseCase) Login(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if !user.CheckPassword(password) {
		return nil, fmt.Errorf("invalid password")
	}

	return user, nil
}
