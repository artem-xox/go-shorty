package service

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/artem-xox/go-shorty/internal/store"
)

type Store interface {
	SetLink(ctx context.Context, link store.Link) error
	GetLink(ctx context.Context, link store.Link) (store.Link, error)
}

type ShortyService struct {
	Store  Store
	Logger *logrus.Logger
}
