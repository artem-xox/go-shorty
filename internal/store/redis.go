package store

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisStorage struct {
	redis *redis.Client
}

func (s *RedisStorage) SetLink(ctx context.Context, link Link) error {
	return nil
}

func (s *RedisStorage) GetLink(ctx context.Context, link Link) (Link, error) {
	return Link{}, nil
}
