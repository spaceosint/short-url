package inmemory

import (
	"errors"
	"github.com/spaceosint/short-url/internal/config"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

type Storage interface {
	GetAll() (map[string]string, error)
	GetOriginalURL(Identifier string) (string, error)
	GetShortURL(cfg config.Config, newUserURL string) (string, error)
}

type UserURL struct {
	Identifier  string `json:"result"`
	OriginalURL string `json:"url"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
}
