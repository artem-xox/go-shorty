package store

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

func TestNewLink(t *testing.T) {
	longLink := "https://www.google.com"
	link := NewLink(longLink)

	if link.Long != longLink {
		t.Errorf("Expected Long to be %q, but got %q", longLink, link.Long)
	}

	if len(link.Hash) != 6 {
		t.Errorf("Expected Hash to be 6 characters, but got %d", len(link.Hash))
	}

	// Verify that the hash is correct
	h := sha1.New()
	h.Write([]byte(longLink))
	expectedHash := fmt.Sprintf("%x", h.Sum(nil))[:6]
	if link.Hash != expectedHash {
		t.Errorf("Expected Hash %q, but got %q", expectedHash, link.Hash)
	}
}
