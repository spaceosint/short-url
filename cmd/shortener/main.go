package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type UserURL struct {
	ID          uint32
	Identifier  string
	OriginalURL string
}

var UsersURL = []UserURL{
	{ID: 1000, OriginalURL: "https://yandex.ru", Identifier: "/t1"},
	{ID: 1001, OriginalURL: "https://yandex.ru/123", Identifier: "/t2"},
}

func getNewID() uint32 {
	max := UsersURL[0].ID
	for _, user := range UsersURL {
		if user.ID > max {
			max = user.ID
		}
	}
	max++
	return max
}
func EndPoint(users []UserURL) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			b, err := io.ReadAll(r.Body)
			// обрабатываем ошибку
			if err != nil {
				http.Error(rw, err.Error(), 500)
				return
			}
			id := getNewID()
			s := ShortenURL(id)

			var newUser = UserURL{ID: id, OriginalURL: string(b), Identifier: "/" + s}
			UsersURL = append(UsersURL, newUser)

			rw.WriteHeader(http.StatusCreated)
			rw.Write([]byte("http://127.0.0.1:8080/" + string(s)))

			//request, err := http.NewRequest("POST", APIURL, nil)
			//if err != nil {
			//	return
			//}

			return
		}
		if r.Method == http.MethodGet {
			uri := r.RequestURI
			for _, user := range UsersURL {
				if user.Identifier == uri {
					rw.Header().Set("Location", user.OriginalURL)
					rw.WriteHeader(http.StatusTemporaryRedirect)
					return
				}
			}
			rw.WriteHeader(http.StatusNoContent)
			return
		} else {
			http.Error(rw, "Only GET or POST requests are allowed!", http.StatusMethodNotAllowed)
			return
		}

	}
}

//func EndPoint(w http.ResponseWriter, r *http.Request) {

func main() {
	// маршрутизация запросов обработчику
	http.HandleFunc("/", EndPoint(UsersURL))
	// запуск сервера с адресом localhost, порт 8080
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// ////////////
const alphabet = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ12345678"

var alphabetLen = uint32(len(alphabet))

func ShortenURL(id uint32) string {
	var (
		nums    []uint32
		num     = id
		builder strings.Builder
	)
	for num > 0 {
		nums = append(nums, num%alphabetLen)
		num /= alphabetLen
	}
	Reverse(nums)

	for _, num := range nums {
		builder.WriteString(string(alphabet[num]))
	}
	return builder.String()
}

func Reverse(s []uint32) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
