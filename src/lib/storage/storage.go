package storage

import (
	"errors"
	"time"
	"context"
)

type Storage interface {
	Saver
	Remover
	Picker
	IsExister
}

type Saver interface {
	SavePage(ctx context.Context, p *Page) error
	SaveNewUser(ctx context.Context, username string) error
	SaveRemind(ctx context.Context, r *Reminds) error
}

type Remover interface {
	RemoveUser(ctx context.Context, userName string) error
	RemovePage(ctx context.Context, p *Page) error
	RemoveRemind(ctx context.Context, r *Reminds) error
}

type Picker interface {
	PickRandomURL(ctx context.Context, userName string) (*Page, error)
	PickReminds(ctx context.Context, userName string) ([]Reminds, error)
}

type IsExister interface {
	IsExistsUser(ctx context.Context, userName string) (bool, error)
	IsExistsURL(ctx context.Context, p *Page) (bool, error)
	IsExistsRemind(ctx context.Context, r *Reminds) (bool, error)
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
