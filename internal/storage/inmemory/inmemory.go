package inmemory

import (
	"github.com/spaceosint/short-url/internal/config"
	"github.com/spaceosint/short-url/internal/storage"
	"github.com/spaceosint/short-url/pkg/shorten"
	"sync"
)

type InMemory struct {
	lock sync.Mutex
	cfg  config.ConfigViper
	m    map[string]string
}

func NewInMemory(config config.ConfigViper) *InMemory {
	return &InMemory{
		cfg: config,
		m:   make(map[string]string),
	}
}

var ID uint = 10000

func (s *InMemory) GetAll() (map[string]string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.m != nil {
		return s.m, nil
	}
	return map[string]string{}, nil
}

func (s *InMemory) GetOriginalURL(Identifier string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if v, ok := s.m[Identifier]; ok {
		return v, nil
	}
	return "", storage.ErrNotFound

}

func (s *InMemory) GetShortURL(newUserURL string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	ID++
	shortURL := shorten.ShortenURL(ID)

	s.m[shortURL] = newUserURL

	return s.cfg.BaseURL + "/" + shortURL, nil
}
