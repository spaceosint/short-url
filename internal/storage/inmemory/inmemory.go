package inmemory

import (
	"fmt"
	"github.com/segmentio/encoding/json"
	"github.com/spaceosint/short-url/internal/config"
	"github.com/spaceosint/short-url/internal/storage"
	"github.com/spaceosint/short-url/pkg/shorten"
	"sync"
)

type InMemory struct {
	lock sync.Mutex
	cfg  config.Config
	m    map[string]string
	//	memory filestore.FileStore
}

func NewInMemory(config config.Config) *InMemory {
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

	// Маршалим данные в JSON
	jsonData, err := json.Marshal(s.m)
	if err != nil {
		fmt.Println(err)

	}

	// Анмаршалим JSON данные в структуру
	var person storage.UserURL
	err = json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println(err)

	}

	// Выводим данные структуры
	fmt.Println(person)
	fmt.Printf("%+v\n", person)
	return s.m, storage.ErrNotFound
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
	s.m["Identifier"] = shortURL
	s.m["OriginalURL"] = newUserURL

	return s.cfg.BaseURL + "/" + shortURL, nil
}
