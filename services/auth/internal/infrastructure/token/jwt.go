package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sssoultrix/event-go/services/auth/internal/application/interfaces"
	"github.com/sssoultrix/event-go/services/auth/internal/config"
	"github.com/sssoultrix/event-go/services/auth/internal/domain"
)

type jwtManager struct {
	cfg config.JWTConfig
}

func NewJWTManager(cfg config.JWTConfig) interfaces.TokenManager {
	return &jwtManager{cfg: cfg}
}

func (m *jwtManager) CreatePair(userID string) (*domain.TokenPair, error) {
	accessToken, err := m.createAccessToken(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	refreshToken, err := m.createRefreshToken(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token: %w", err)
	}

	return &domain.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (m *jwtManager) Verify(tokenStr string) (*domain.AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.cfg.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if claims, ok := token.Claims.(*domain.AccessTokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func (m *jwtManager) createAccessToken(userID string) (string, error) {
	claims := &domain.AccessTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.cfg.AccessTokenTTL)),
			ID:        uuid.NewString(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.cfg.SecretKey))
}

func (m *jwtManager) createRefreshToken(userID string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.cfg.RefreshTokenTTL)),
		ID:        uuid.NewString(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.cfg.SecretKey))
}
