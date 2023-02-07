package storage

import (
	"sync"
)

type InMemory struct {
	lock sync.Mutex
	m    map[string]string
}

func NewInMemory() *InMemory {
	return &InMemory{
		m: make(map[string]string),
	}
}

var ID int = 10000

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

func (s *InMemory) GetShortURL(newUserURL string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	ID++
	shortURL := ShortenURL(ID)

	if _, ok := s.m[shortURL]; ok {
		return "", ErrAlreadyExists
	}
	s.m[shortURL] = newUserURL

	return shortURL, nil
}

