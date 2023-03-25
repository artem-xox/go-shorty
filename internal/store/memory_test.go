package store

import (
	"context"
	"testing"
)

func TestMemoryCacheStore(t *testing.T) {
	ctx := context.Background()

	s := NewMemoryCacheStore()

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
