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
	Updater
}

type Saver interface {
	SavePage(ctx context.Context, p *Page) error
	SaveNewUser(ctx context.Context, username string) error
	// SaveRemind(ctx context.Context, r *Reminds) error
}

type Remover interface {
	RemoveUser(ctx context.Context, userName string) error
	RemovePage(ctx context.Context, p *Page) error
	// RemoveRemind(ctx context.Context, r *Reminds) error
}

type Picker interface {
	PickRandomPage(ctx context.Context, userName string) (*Page, error)
	PickUserInfo(ctx context.Context, userName string) (*UserInfo, error)
	// PickReminds(ctx context.Context, userName string) ([]Reminds, error)
}

type IsExister interface {
	IsExistsUser(ctx context.Context, userName string) (bool, error)
	IsExistsPage(ctx context.Context, p *Page) (bool, error)
	// IsExistsRemind(ctx context.Context, r *Reminds) (bool, error)
}

type Updater interface {
	UpdateUserInfo(ctx context.Context, username string, messageID int, chatID int) error
}

type Page struct {
	URL string
	UserName string
}

type UserInfo struct {
	ChatID int
	MessageID int
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
