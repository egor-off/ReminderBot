package sqlite

import (
	"context"
	"database/sql"
	"src/lib/e"
	"src/lib/storage"
)

const (
	ErrCannotOpenDB = "cannot open DB"
	ErrCannotCreateTable = "cannot create DB"
	ErrCannotPingDB = "cannot connect to DB"
	ErrCannotSavePage = "cannot save page to DB"
	ErrSelectFromDB = "cannot select data from DB"
	ErrDeleteFromDB = "cannot delete data from DB"
	ErrIsExists = "cannot check if exists the page"
	ErrAddNewUser = "cannot add new user"
)

type Storage struct {
	db *sql.DB
}

// New open database by path and return *Storage.
func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, e.Wrap(ErrCannotOpenDB, err)
	}

	if err := db.Ping(); err != nil {
		return nil, e.Wrap(ErrCannotPingDB, err)
	}

	return &Storage{db: db}, nil
}

//SaveNewUser func saves new user to database.
func (s *Storage) SaveNewUser(ctx context.Context, username string) error {
	if _, err := s.db.ExecContext(ctx, insertNewUser, username); err != nil {
		return e.Wrap(ErrAddNewUser, err)
	}
	return nil
}

// Save func saves page to database.
func (s *Storage) SavePage(ctx context.Context, p *storage.Page) error {
	if _, err := s.db.ExecContext(ctx, insertURL, p.UserName, p.URL); err != nil {
		return e.Wrap(ErrCannotSavePage, err)
	}

	return nil
}

// SaveRemind saves remind to database.
func (s *Storage) SaveRemind(ctx context.Context, p *storage.Reminds) error {
	if _, err := s.db.ExecContext(ctx, insertRemind, p.UserName, p.Message, p.Date, p.Period); err != nil {
		return e.Wrap("cannot save remind", err)
	}
	return nil
}

// PickRandom gives random URL from database by username.
func (s *Storage) PickRandom(ctx context.Context, username string) (*storage.Page, error) {
	var url string

	if err := s.db.QueryRowContext(ctx, pickRandom, username).Scan(&url); err == sql.ErrNoRows {
		return nil, storage.ErrNoSavedPages
	} else if err != nil {
		return nil, e.Wrap(ErrSelectFromDB, err)
	}

	return &storage.Page{URL: url, UserName: username}, nil
}

// Remove delete page from database.
func (s *Storage) Remove(ctx context.Context, p *storage.Page) error {
	q := `DELETE FROM pages WHERE url = ? AND user_name = ?`
	if _, err := s.db.ExecContext(ctx, q, p.URL, p.UserName); err != nil {
		return e.Wrap(ErrDeleteFromDB, err)
	}
	return nil
}

// IsExists check if storage exists.
func (s *Storage) IsExists(ctx context.Context, p *storage.Page) (bool, error) {
	q := `SELECT COUNT(*) FROM pages WHERE url = ? and user_name = ?`

	var count int

	if err := s.db.QueryRowContext(ctx, q, p.URL, p.UserName).Scan(&count); err != nil {
		return false, e.Wrap(ErrIsExists, err)
	}

	return count > 0, nil
}

func (s *Storage) Init(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, createTable)
	if err != nil {
		return e.Wrap(ErrCannotCreateTable, err)
	}

	return nil
}
