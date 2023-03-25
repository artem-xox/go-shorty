package store

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestRedisCacheStore(t *testing.T) {
	ctx := context.Background()

	// Connect to a Redis server
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	s := RedisCacheStore{
		client: rdb,
	}

	// Add a new link to the cache
	longLink := "https://www.github.com"
	link := NewLink(longLink)
	err := s.SetLink(ctx, link)
	if err != nil {
		t.Fatalf("Error setting link: %v", err)
	}

	// Attempt to retrieve a link that's not in the cache
	newLongLink := "https://www.google.com"
	newLink := NewLink(newLongLink)
	_, err = s.GetLink(ctx, newLink)
	if err == nil {
		t.Fatalf("Expected link not to be found")
	}

	// Retrieve a link that's in the cache
	cachedLink, err := s.GetLink(ctx, link)
	if err != nil {
		t.Fatalf("Error retrieving link from cache: %v", err)
	}
	if cachedLink.Long != longLink {
		t.Fatalf("Expected Long link to match")
	}
	if cachedLink.Hash != link.Hash {
		t.Fatalf("Expected Hash link to match")
	}
}
