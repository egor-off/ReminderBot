package storage

import (
	"context"
	"database/sql"
	"log"
	"src/lib/e"
	"src/lib/storage"

	_ "github.com/mattn/go-sqlite3"
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
	ErrIsRemovedURL
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

//SaveNewUser saves new user to database.
func (s *Storage) SaveNewUser(ctx context.Context, username string) error {
	if _, err := s.db.ExecContext(ctx, insertNewUser, username); err != nil {
		return e.Wrap(ErrAddNewUser, err)
	}
	return nil
}

// SavePage saves page to database.
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

func (s *Storage) UpdateUserInfo(ctx context.Context, username string, messageID int, chatID int) error {
	if _, err := s.db.ExecContext(ctx, updateUsersIDs, messageID, chatID, username); err != nil {
		return e.Wrap("cannot write IDs: ", err)
	}
	return nil
}

func (s *Storage) PickUserInfo(ctx context.Context, username string) (*storage.UserInfo, error) {
	u := &storage.UserInfo{}

	if err := s.db.QueryRowContext(ctx, pickUserInfo, username).Scan(&(u.MessageID), &(u.ChatID)); err != nil {
		return nil, e.Wrap("cannot get userinfo: ", err)
	}

	log.Printf("userinfo: %v", u)

	return u, nil
}

// PickRandomPage gives random URL from database by username.
func (s *Storage) PickRandomPage(ctx context.Context, username string) (*storage.Page, error) {
	var url string

	if err := s.db.QueryRowContext(ctx, pickRandomURL, username).Scan(&url); err == sql.ErrNoRows {
		return nil, storage.ErrNoSavedPages
	} else if err != nil {
		return nil, e.Wrap(ErrSelectFromDB, err)
	}

	return &storage.Page{URL: url, UserName: username}, nil
}

// Remove delete page from database.
func (s *Storage) RemovePage(ctx context.Context, p *storage.Page) error {
	// q := `DELETE FROM pages WHERE url = ? AND user_name = ?`
	if _, err := s.db.ExecContext(ctx, deleteURL, p.URL, p.UserName); err != nil {
		return e.Wrap(ErrDeleteFromDB, err)
	}
	return nil
}

// RemoveUser delete user from database.
func (s *Storage) RemoveUser(ctx context.Context, userName string) error {
	if _, err := s.db.ExecContext(ctx, deleteUser, userName); err != nil {
		return e.Wrap(ErrDeleteFromDB, err)
	}
	return nil
}

// IsExists check if storage exists.
func (s *Storage) IsExistsPage(ctx context.Context, p *storage.Page) (bool, error) {
	var count int

	if err := s.db.QueryRowContext(ctx, isExistsURL, p.URL, p.UserName).Scan(&count); err != nil {
		return false, e.Wrap(ErrIsExists, err)
	}

	return count > 0, nil
}

//IsExistsUser check if user allready exists.
func (s *Storage) IsExistsUser(ctx context.Context, userName string) (bool, error) {
	var count int

	if err := s.db.QueryRowContext(ctx, isExistsUser, userName).Scan(&count); err != nil {
		return false, e.Wrap(ErrIsExists, err)
	}

	return count > 0, nil
}

func (s *Storage) IsRemovedURL(ctx context.Context, text string, username string) (bool, error) {
	var count int

	if err := s.db.QueryRowContext(ctx, isRemovedURL, text, username).Scan(&count); err != nil {
		return false, e.Wrap(ErrIsRemovedURL, err)
	}
	return count > 0, nil
}

func (s *Storage) UpdateURLRemoved(ctx context.Context, text string, username string) error {
	if _, err := s.db.ExecContext(ctx, updateURLRemoved, text, username); err != nil {
		return e.Wrap("cannot rewrite removed: ", err)
	}
	return nil
}

// Init creates tables for database.
func (s *Storage) Init(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, createTables)
	if err != nil {
		return e.Wrap(ErrCannotCreateTable, err)
	}

	return nil
}
