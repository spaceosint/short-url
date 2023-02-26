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
	GetAll() ([]MyStruct, error)
	GetOriginalURL(cfg config.Config, Identifier string) (string, error)
	GetShortURL(uuid any, cfg config.Config, newUserURL string) (string, error)
	GetAllByCookie(cfg config.Config, uuid any) ([]respData, error)
}

type UserURL struct {
	Identifier  string `json:"result"`
	OriginalURL string `json:"url"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
}
