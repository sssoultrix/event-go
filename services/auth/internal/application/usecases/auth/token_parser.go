package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func (uc *authUseCase) parseRefreshToken(refreshToken string) (*jwt.RegisteredClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(refreshToken, &jwt.RegisteredClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse refresh token: %w", err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("invalid refresh token claims")
	}

	return claims, nil
}
