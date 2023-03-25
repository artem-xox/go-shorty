package store

import (
	"context"
)

type MemoryCacheStore struct {
	cache map[string]Link
}

func NewMemoryCacheStore() *MemoryCacheStore {
	return &MemoryCacheStore{
		cache: make(map[string]Link),
	}
}

func (s *MemoryCacheStore) SetLink(ctx context.Context, link Link) error {
	s.cache[link.Hash] = link
	return nil
}

func (s *MemoryCacheStore) GetLink(ctx context.Context, link Link) (Link, error) {
	cachedLink, ok := s.cache[link.Hash]
	if ok {
		return cachedLink, nil
	}

	// link not in cache, return an error
	return Link{}, ErrorLinkNotFound
}
