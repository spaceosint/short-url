package storage

import (
	"sync"
)

type InMemory struct {
	lock sync.Mutex
	m    map[string]string
}

func NewInMemory() *InMemory {
	return &InMemory{}
}

var UsersURL = []UserURL{
	{ID: 1000, OriginalURL: "https://yandex.ru", Identifier: "t1"},
	{ID: 1001, OriginalURL: "https://yandex.ru/123", Identifier: "t2"},
}

func (s *InMemory) GetAll() ([]UserURL, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if UsersURL != nil {
		return UsersURL, nil
	}
	return []UserURL{}, ErrNotFound
}
func (s *InMemory) GetOriginalURL(Identifier string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, user := range UsersURL {
		if user.Identifier == Identifier {
			return user.OriginalURL, nil
		}
	}
	return "", ErrNotFound
}

func (s *InMemory) GetShortURL(newUserURL string) string {
	s.lock.Lock()
	defer s.lock.Unlock()
	id := GetNewID()
	shortURL := ShortenURL(id)
	var newUser = UserURL{ID: id, OriginalURL: newUserURL, Identifier: shortURL}
	UsersURL = append(UsersURL, newUser)
	return shortURL
}

func GetNewID() uint32 {
	max := UsersURL[0].ID
	for _, user := range UsersURL {
		if user.ID > max {
			max = user.ID
		}
	}
	max++
	return max
}
