package storage

import "errors"

type UserURL struct {
	ID          uint32
	Identifier  string
	OriginalURL string
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

var (
	ErrNotFound = errors.New("not found")
	//ErrAlreadyExists = errors.New("already exists")
)

type Storage interface {
	GetAll() (map[string]map[string]string, error)
	GetOriginalURL(Identifier string) (string, error)
	GetShortURL(newUserURL string) string
}
