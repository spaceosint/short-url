package handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/spaceosint/short-url/internal/storage"
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
				expectedResponseBody: "http://127.0.0.1:8080/dkK",
			}}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				// Init Endpoint
				r := gin.New()

				st := storage.NewInMemory()

				r.POST("/", New(st).PostNewUserURL)

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
		type request struct {
			inputBody string
		}
		tests := []struct {
			name                   string
			getParams              string
			expectedStatusCode     int
			expectedResponseHeader string
			newRequest             request
		}{
			{
				name:                   "Ok GET",
				getParams:              "dkK",
				expectedStatusCode:     307,
				expectedResponseHeader: "https://google.com",
				newRequest:             request{inputBody: "https://google.com"},
			}}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				// Init Endpoint
				r := gin.New()
				st := storage.NewInMemory()

				r.POST("/", New(st).PostNewUserURL)
				// Create Request
				w := httptest.NewRecorder()
				nReq := httptest.NewRequest("POST", "/",
					bytes.NewBufferString(test.newRequest.inputBody))

				// Make Request
				r.ServeHTTP(w, nReq)
				r2 := gin.New()
				r2.GET("/:Identifier", New(st).GetUserURLByIdentifier)
				// Create Request
				w2 := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "/"+test.getParams, nil)

				// Make Request
				r2.ServeHTTP(w2, req)

				// Assert
				assert.Equal(t, test.expectedStatusCode, w2.Code)
				assert.Equal(t, test.expectedResponseHeader, w2.Header().Get("Location"))
			})
		}
	})
}
func TestHandler_PostNewUserURLJSON(t *testing.T) {
	t.Run("handler", func(t *testing.T) {
		tests := []struct {
			name                 string
			inputBody            string
			expectedStatusCode   int
			expectedResponseBody string
		}{
			{
				name:               "Ok POST",
				inputBody:          `{"url": "https://google.com/new2"}`,
				expectedStatusCode: 201,
				expectedResponseBody: `{
    "result": "http://127.0.0.1:8080/dkK"
}`,
			}}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				// Init Endpoint
				r := gin.New()

				st := storage.NewInMemory()

				r.POST("/api/shorten", New(st).PostNewUserURLJSON)

				// Create Request
				m, b := map[string]string{"url": "https://google.com/new2"}, new(bytes.Buffer)
				json.NewEncoder(b).Encode(m)

				w := httptest.NewRecorder()
				req := httptest.NewRequest("POST", "/api/shorten",
					b)

				// Make Request
				r.ServeHTTP(w, req)

				// Assert
				assert.Equal(t, w.Code, test.expectedStatusCode)
				assert.Equal(t, w.Body.String(), test.expectedResponseBody)
			})
		}
	})
}
