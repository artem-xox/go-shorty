package store

import (
	"crypto/sha1"
	"fmt"
)

type Link struct {
	Long string
	// sixth symbol hash
	Hash string
}

func NewLink(long string) Link {
	// Generate a hash from the long link
	h := sha1.New()
	h.Write([]byte(long))
	bs := h.Sum(nil)
	hash := fmt.Sprintf("%x", bs)[:6] // Cut hash to 6 characters

	return Link{
		Long: long,
		Hash: hash,
	}
}
