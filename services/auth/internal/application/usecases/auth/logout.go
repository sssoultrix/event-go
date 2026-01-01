package auth

import (
	"context"
	"fmt"
)

func (uc *authUseCase) Logout(ctx context.Context, refreshToken string) error {
	claims, err := uc.parseRefreshToken(refreshToken)
	if err != nil {
		return err
	}

	if err := uc.tokenStore.Delete(ctx, claims.Subject, claims.ID); err != nil {
		return fmt.Errorf("failed to delete refresh token: %w", err)
	}

	return nil
}
