package inmemory

import (
	"fmt"
	"github.com/spaceosint/short-url/internal/config"
	"github.com/spaceosint/short-url/internal/storage/filestore"
	"github.com/spaceosint/short-url/pkg/shorten"
	"sync"
)

type InMemory struct {
	lock   sync.Mutex
	data   MyStruct
	memory filestore.FileStore
}

func NewInMemory() *InMemory {
	return &InMemory{}
}

type MyStruct struct {
	UUID string
	Data respData
}
type respData struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

var MySlice []MyStruct
var ID uint = 10000

func (s *InMemory) GetAll() ([]MyStruct, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	fmt.Println("MySlice")
	fmt.Println(MySlice)
	if MySlice != nil {
		return MySlice, nil
	}
	return []MyStruct{}, ErrNotFound
}
func (s *InMemory) GetOriginalURL(Identifier string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for id := range MySlice {
		if MySlice[id].Data.ShortURL == Identifier {
			return MySlice[id].Data.OriginalURL, nil
		}
	}

	return "", ErrNotFound

}

func (s *InMemory) GetShortURL(uuid any, cfg config.Config, newUserURL string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	ID++
	shortURL := shorten.ShortenURL(ID)

	var newUser = MyStruct{UUID: uuid.(string), Data: respData{ShortURL: cfg.BaseURL + "/" + shortURL, OriginalURL: newUserURL}}
	MySlice = append(MySlice, newUser)

	return cfg.BaseURL + "/" + shortURL, nil
}
func (s *InMemory) GetAllByCookie(cfg config.Config, uuid any) ([]respData, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	var resp []respData
	if MySlice != nil {
		for id := range MySlice {
			if MySlice[id].UUID == uuid.(string) {

				resp = append(resp, MySlice[id].Data)
			}
		}
		return resp, nil

	}
	return []respData{}, ErrNotFound
}
