package storage

import (
	"fmt"
	"strconv"
	"sync"
)

type InMemory struct {
	lock sync.Mutex
	m    map[string]map[string]string
}

func NewInMemory() *InMemory {
	return &InMemory{
		m: make(map[string]map[string]string),
	}
	//	m: make(map[int]string),
	//}
}

//MapUserURL := make(map[int]string)

//var UsersURL = []UserURL{
//	{ID: 1000, OriginalURL: "https://yandex.ru", Identifier: "t1"},
//	{ID: 1001, OriginalURL: "https://yandex.ru/123", Identifier: "t2"},
//}

func (s *InMemory) GetAll() (map[string]map[string]string, error) {
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

	for _, userUrls := range s.m {
		for shortUserUrl, originalUrl := range userUrls {
			if Identifier == shortUserUrl {
				return originalUrl, nil
			}
		}
	}
	//if user.Identifier == Identifier {
	//	return user.OriginalURL, nil
	//}

	return "", ErrNotFound
}

func (s *InMemory) GetShortURL(newUserURL string) string {
	s.lock.Lock()
	defer s.lock.Unlock()
	fmt.Println(newUserURL)
	id := s.GetNewID()

	shortURL := ShortenURL(id)
	fmt.Println(id, shortURL)
	s.m[strconv.Itoa(id)] = map[string]string{
		shortURL: newUserURL,
	}
	fmt.Println(s.m)
	//var newUser = UserURL{ID: id, OriginalURL: newUserURL, Identifier: shortURL}
	//UsersURL = append(UsersURL, newUser)
	return shortURL
}

func (s *InMemory) GetNewID() int {
	var max = 10001
	for stId, _ := range s.m {
		id, _ := strconv.Atoi(stId)
		if id > max {
			max = id
		}
	}
	max++
	return max
}
