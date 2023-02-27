package inmemory

import (
	"github.com/spaceosint/short-url/internal/config"
	"github.com/spaceosint/short-url/internal/storage/filestore"
	"github.com/spaceosint/short-url/pkg/shorten"
	"sync"
)

type InMemory struct {
	lock   sync.Mutex
	m      map[string]string
	memory filestore.FileStore
}

func NewInMemory() *InMemory {
	return &InMemory{
		m: make(map[string]string),
	}
}

var ID uint = 10000

func (s *InMemory) GetAll() (map[string]string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.m != nil {
		return s.m, nil
	}
	return s.m, ErrNotFound
}
func (s *InMemory) GetOriginalURL(Identifier string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if v, ok := s.m[Identifier]; ok {
		return v, nil
	}
	return "", ErrNotFound

}

func (s *InMemory) GetShortURL(cfg config.Config, newUserURL string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	ID++
	shortURL := shorten.ShortenURL(ID)

	s.m[shortURL] = newUserURL

	return cfg.BaseURL + "/" + shortURL, nil
}
