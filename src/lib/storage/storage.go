package storage

import (
	"errors"
	"time"
	"context"
)

type Storage interface {
	SavePage(ctx context.Context, p *Page) error
	PickRandom(ctx context.Context, userName string) (*Page, error)
	Remove(ctx context.Context, p *Page) error
	IsExistsURL(ctx context.Context, p *Page) (bool, error)
	// work with reminds needed
}

type Page struct {
	URL string
	UserName string
}

type Reminds struct {
	UserName string
	Message string
	Date time.Time
	Period time.Time
}

var (
	ErrNoSavedPages = errors.New("no saved pages")
)
