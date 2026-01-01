package auth

import (
	"context"
	"fmt"

	"github.com/sssoultrix/event-go/services/auth/internal/domain"
)

func (uc *authUseCase) Login(ctx context.Context, params domain.LoginParams) (*domain.TokenPair, error) {
	user, err := uc.usersService.Login(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("login failed: %w", err)
	}

	tokenPair, err := uc.tokenManager.CreatePair(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create token pair: %w", err)
	}

	claims, err := uc.parseRefreshToken(tokenPair.RefreshToken)
	if err != nil {
		return nil, err
	}

	expiresIn := claims.ExpiresAt.Sub(claims.IssuedAt.Time)
	if err := uc.tokenStore.Store(ctx, user.ID, claims.ID, expiresIn); err != nil {
		return nil, fmt.Errorf("failed to store refresh token: %w", err)
	}

	return tokenPair, nil
}
