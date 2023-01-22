package handlers

import (
	"ShortURL/internal/app/shorten"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_PostNewUserURL(t *testing.T) {
	t.Run("handler", func(t *testing.T) {
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
				r := gin.New()
				s := shorten.New()
				r.POST("/", New(s).PostNewUserURL)

				// Create Request
				w := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "/",
					bytes.NewBufferString(test.inputBody))

				// Make Request
				r.ServeHTTP(w, req)

				// Assert
				assert.Equal(t, w.Code, test.expectedStatusCode)
				assert.Equal(t, w.Body.String(), test.expectedResponseBody)
			})
		}
	})
}

func TestHandler_GetUserURLByIdentifier(t *testing.T) {
	t.Run("handler", func(t *testing.T) {
		tests := []struct {
			name                   string
			getParams              string
			expectedStatusCode     int
			expectedResponseHeader string
		}{
			{
				name:                   "Ok GET",
				getParams:              "t1",
				expectedStatusCode:     307,
				expectedResponseHeader: "https://yandex.ru",
			}}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				// Init Endpoint
				r := gin.New()
				s := shorten.New()
				r.GET("/:Identifier", New(s).GetUserURLByIdentifier)

				// Create Request
				w := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "/"+test.getParams, nil)

				// Make Request
				r.ServeHTTP(w, req)

				// Assert
				assert.Equal(t, test.expectedStatusCode, w.Code)
				assert.Equal(t, test.expectedResponseHeader, w.HeaderMap.Get("Location"))
			})
		}
	})
}
