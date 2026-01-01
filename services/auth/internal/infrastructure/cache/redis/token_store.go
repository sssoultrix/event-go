package redis

import (
	"context"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/sssoultrix/event-go/services/auth/internal/application/interfaces"
)

type TokenStore struct {
	cache *Cache
}

func NewTokenStore(client *goredis.Client) interfaces.TokenStore {
	return &TokenStore{cache: NewCache(client)}
}

func (s *TokenStore) Store(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error {
	key := s.key(userID, tokenID)
	return s.cache.Set(ctx, key, []byte("1"), expiresIn)
}

func (s *TokenStore) Get(ctx context.Context, userID string, tokenID string) (bool, error) {
	key := s.key(userID, tokenID)
	_, found, err := s.cache.Get(ctx, key)
	return found, err
}

func (s *TokenStore) Delete(ctx context.Context, userID string, tokenID string) error {
	key := s.key(userID, tokenID)
	return s.cache.Del(ctx, key)
}

func (s *TokenStore) key(userID, tokenID string) string {
	return fmt.Sprintf("user:%s:token:%s", userID, tokenID)
}
