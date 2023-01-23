package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	var TestUsersURL = []UserURL{
		{ID: 1000, OriginalURL: "https://yandex.ru", Identifier: "/t1"},
		{ID: 1001, OriginalURL: "https://yandex.ru/123", Identifier: "/t2"},
	}

	t.Run("POST", func(t *testing.T) {

		tests := []struct {
			name                 string
			inputBody            string
			expectedStatusCode   int
			expectedResponseBody string
		}{
			{
				name:                 "Ok POST",
				inputBody:            "https://google.com",
				expectedStatusCode:   201,
				expectedResponseBody: "http://127.0.0.1:8080/t3",
			}}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				// Init Endpoint
				request := httptest.NewRequest(http.MethodPost, "/", nil)
				w := httptest.NewRecorder()

				h := http.HandlerFunc(EndPoint(TestUsersURL))
				h(w, request)
				result := w.Result()
				assert.Equal(t, test.expectedStatusCode, result.StatusCode)
				assert.Equal(t, test.expectedResponseBody, w.Body.String())

			})
		}

	})
	t.Run("GET", func(t *testing.T) {

		tests := []struct {
			name                 string
			inputBody            string
			expectedStatusCode   int
			expectedResponseBody string
		}{
			{
				name:                 "Ok GET",
				inputBody:            "/t2",
				expectedStatusCode:   307,
				expectedResponseBody: "https://yandex.ru/123",
			}}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				// Init Endpoint
				request := httptest.NewRequest(http.MethodGet, test.inputBody, nil)
				w := httptest.NewRecorder()

				h := http.HandlerFunc(EndPoint(TestUsersURL))
				h(w, request)
				result := w.Result()
				assert.Equal(t, test.expectedStatusCode, result.StatusCode)
				assert.Equal(t, test.expectedResponseBody, w.HeaderMap.Get("Location"))

			})
		}

	})
}
