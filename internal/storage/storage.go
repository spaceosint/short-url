package storage

import (
	"errors"
	"github.com/spaceosint/short-url/internal/config"
	"github.com/spaceosint/short-url/internal/storage/filestore"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

type Storage interface {
	GetAll(config config.Config) ([]filestore.Event, error)
	GetOriginalURL(config config.Config, Identifier string) (string, error)
	GetShortURL(config config.Config, newUserURL string) (string, error)
}

type UserURL struct {
	Identifier  string `json:"result"`
	OriginalURL string `json:"url"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
}
