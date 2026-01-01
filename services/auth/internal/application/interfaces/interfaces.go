package interfaces

import (
	"context"
	"time"

	"github.com/sssoultrix/event-go/services/auth/internal/domain"
)

type TokenManager interface {
	CreatePair(userID string) (*domain.TokenPair, error)
	Verify(token string) (*domain.AccessTokenClaims, error)
}

type TokenStore interface {
	Store(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error
	Get(ctx context.Context, userID string, tokenID string) (bool, error)
	Delete(ctx context.Context, userID string, tokenID string) error
}
