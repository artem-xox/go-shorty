package store

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type RedisCacheStore struct {
	client *redis.Client
}

func NewRedisCacheStore(client *redis.Client) *RedisCacheStore {
	return &RedisCacheStore{client: client}
}

func (s *RedisCacheStore) SetLink(ctx context.Context, link Link) error {
	json, err := json.Marshal(link)
	if err != nil {
		return err
	}
	err = s.client.Set(ctx, link.Hash, json, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *RedisCacheStore) GetLink(ctx context.Context, link Link) (Link, error) {
	val, err := s.client.Get(ctx, link.Hash).Result()
	if err == redis.Nil {
		// link not in cache, return an error
		return Link{}, ErrorLinkNotFound
	} else if err != nil {
		// redis error
		return Link{}, err
	}
	var cachedLink Link
	err = json.Unmarshal([]byte(val), &cachedLink)
	if err != nil {
		return Link{}, err
	}
	return cachedLink, nil
}
