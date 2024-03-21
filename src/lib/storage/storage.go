package storage

import (
	"errors"
	"time"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
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
