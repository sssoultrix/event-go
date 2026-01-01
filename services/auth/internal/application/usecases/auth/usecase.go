package auth

import (
	"context"

	"github.com/sssoultrix/event-go/services/auth/internal/application/interfaces"
	"github.com/sssoultrix/event-go/services/auth/internal/domain"
)

type AuthUseCase interface {
	Login(ctx context.Context, params domain.LoginParams) (*domain.TokenPair, error)
	Refresh(ctx context.Context, refreshToken string) (*domain.TokenPair, error)
	Logout(ctx context.Context, refreshToken string) error
}

type authUseCase struct {
	usersService domain.UsersService
	tokenManager interfaces.TokenManager
	tokenStore   interfaces.TokenStore
}

func NewAuthUseCase(us domain.UsersService, tm interfaces.TokenManager, ts interfaces.TokenStore) AuthUseCase {
	return &authUseCase{
		usersService: us,
		tokenManager: tm,
		tokenStore:   ts,
	}
}
