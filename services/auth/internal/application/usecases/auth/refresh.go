package auth

import (
	"context"
	"fmt"

	"github.com/sssoultrix/event-go/services/auth/internal/domain"
)

func (uc *authUseCase) Refresh(ctx context.Context, refreshToken string) (*domain.TokenPair, error) {
	claims, err := uc.parseRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	userID := claims.Subject
	tokenID := claims.ID

	found, err := uc.tokenStore.Get(ctx, userID, tokenID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify refresh token from store: %w", err)
	}
	if !found {
		return nil, fmt.Errorf("refresh token not found or revoked")
	}

	if err := uc.tokenStore.Delete(ctx, userID, tokenID); err != nil {
		return nil, fmt.Errorf("failed to revoke old refresh token: %w", err)
	}

	newTokenPair, err := uc.tokenManager.CreatePair(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to create new token pair: %w", err)
	}

	newClaims, err := uc.parseRefreshToken(newTokenPair.RefreshToken)
	if err != nil {
		return nil, err
	}

	newExpiresIn := newClaims.ExpiresAt.Sub(newClaims.IssuedAt.Time)
	if err := uc.tokenStore.Store(ctx, userID, newClaims.ID, newExpiresIn); err != nil {
		return nil, fmt.Errorf("failed to store new refresh token: %w", err)
	}

	return newTokenPair, nil
}
