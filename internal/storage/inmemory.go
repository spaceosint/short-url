package storage

import (
	"fmt"
	"github.com/spaceosint/short-url/internal/config"
	"github.com/spaceosint/short-url/internal/storage/filestore"
	"sync"
)

type InMemory struct {
	lock   sync.Mutex
	memory filestore.FileStore
}

func NewInMemory() *InMemory {
	return &InMemory{}
}

func (s *InMemory) GetAll(cfg config.Config) ([]filestore.Event, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	//bd, _ := file_bd.NewConsumer("bd")
	//fmt.Println(bd)

	users := s.memory.GetAllByPath(cfg.FileStoragePath)
	if users != nil {
		return users, nil
	}
	return users, ErrNotFound
}
func (s *InMemory) GetOriginalURL(cfg config.Config, Identifier string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	originalURL, err := s.memory.GetOriginalURL(Identifier, cfg.FileStoragePath)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}

func (s *InMemory) GetShortURL(cfg config.Config, newUserURL string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	newID := s.memory.GetNewID(cfg.FileStoragePath)

	shortURL := ShortenURL(newID)
	fmt.Println()
	err := s.memory.AddNewLink(newID, shortURL, newUserURL, cfg.FileStoragePath)
	if err != nil {
		panic(err)
	}

	return cfg.BaseURL + "/" + shortURL, nil
}
