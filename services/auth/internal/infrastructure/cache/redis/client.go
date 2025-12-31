package redis

import (
	"context"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/sssoultrix/event-go/services/auth/internal/config"
)

type Client struct {
	*goredis.Client
}

func NewClient(cfg config.RedisConfig) *Client {
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     cfg.Addr(),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return &Client{Client: rdb}
}

func (c *Client) Ping(ctx context.Context, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return c.Client.Ping(ctx).Err()
}
