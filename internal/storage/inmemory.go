package storage

import (
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

func (s *InMemory) GetAll() ([]filestore.Event, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	//bd, _ := file_bd.NewConsumer("bd")
	//fmt.Println(bd)
	cfg := config.GetConfig()
	users := s.memory.GetAllByPath(cfg.FileStoragePath)
	if users != nil {
		return users, nil
	}
	return users, ErrNotFound
}
func (s *InMemory) GetOriginalURL(Identifier string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	cfg := config.GetConfig()
	originalURL, err := s.memory.GetOriginalURL(Identifier, cfg.FileStoragePath)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}

func (s *InMemory) GetShortURL(newUserURL string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	cfg := config.GetConfig()

	newID := s.memory.GetNewID(cfg.FileStoragePath)

	shortURL := ShortenURL(newID)

	err := s.memory.AddNewLink(newID, shortURL, newUserURL, cfg.FileStoragePath)
	if err != nil {
		panic(err)
	}

	return cfg.BaseURL + "/" + shortURL, nil
}
