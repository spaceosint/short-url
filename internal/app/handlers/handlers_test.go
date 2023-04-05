// package handlers
//
// import (
//
//	"bytes"
//	"github.com/gin-gonic/gin"
//	"github.com/spaceosint/short-url/internal/config"
//	"github.com/spaceosint/short-url/internal/storage/inmemory"
//	"github.com/stretchr/testify/assert"
//	"net/http/httptest"
//	"testing"
//
// )
//
//	func TestHandler_PostNewUserURL(t *testing.T) {
//		t.Run("handler", func(t *testing.T) {
//			tests := []struct {
//				name                 string
//				inputBody            string
//				expectedStatusCode   int
//				expectedResponseBody string
//			}{
//				{
//					name:                 "Ok POST",
//					inputBody:            "https://google.com",
//					expectedStatusCode:   201,
//					expectedResponseBody: "http://127.0.0.1:8080/dkK",
//				}}
//			for _, test := range tests {
//				t.Run(test.name, func(t *testing.T) {
//					// Init Endpoint
//					r := gin.New()
//
//					st := inmemory.NewInMemory()
//					cfg := config.GetConfig()
//					r.POST("/", New(st, cfg).PostNewUserURL)
//
//					// Create Request
//					w := httptest.NewRecorder()
//					req := httptest.NewRequest("POST", "/",
//						bytes.NewBufferString(test.inputBody))
//
//					// Make Request
//					r.ServeHTTP(w, req)
//
//					// Assert
//					assert.Equal(t, w.Code, test.expectedStatusCode)
//					assert.Equal(t, w.Body.String(), test.expectedResponseBody)
//				})
//			}
//		})
//	}
//
// //func TestHandler_GetUserURLByIdentifier(t *testing.T) {
// //	t.Run("handler", func(t *testing.T) {
// //		type request struct {
// //			inputBody string
// //		}
// //		tests := []struct {
// //			name                   string
// //			getParams              string
// //			expectedStatusCode     int
// //			expectedResponseHeader string
// //			newRequest             request
// //		}{
// //			{
// //				name:                   "Ok GET",
// //				getParams:              "dkK",
// //				expectedStatusCode:     307,
// //				expectedResponseHeader: "https://google.com",
// //				newRequest:             request{inputBody: "https://google.com"},
// //			}}
// //		for _, test := range tests {
// //			t.Run(test.name, func(t *testing.T) {
// //				// Init Endpoint
// //				r := gin.New()
// //				st := storage.NewInMemory()
// //
// //				r.POST("/", New(st).PostNewUserURL)
// //				// Create Request
// //				w := httptest.NewRecorder()
// //				nReq := httptest.NewRequest("POST", "/",
// //					bytes.NewBufferString(test.newRequest.inputBody))
// //
// //				// Make Request
// //				r.ServeHTTP(w, nReq)
// //				r2 := gin.New()
// //				r2.GET("/:Identifier", New(st).GetUserURLByIdentifier)
// //				// Create Request
// //				w2 := httptest.NewRecorder()
// //				req := httptest.NewRequest("GET", "/"+test.getParams, nil)
// //
// //				// Make Request
// //				r2.ServeHTTP(w2, req)
// //
// //				// Assert
// //				assert.Equal(t, test.expectedStatusCode, w2.Code)
// //				assert.Equal(t, test.expectedResponseHeader, w2.Header().Get("Location"))
// //			})
// //		}
// //	})
// //}
//
// //func TestHandler_GetUserURLByIdentifier(t *testing.T) {
// //	t.Run("handler", func(t *testing.T) {
// //		type request struct {
// //			inputBody string
// //		}
// //		tests := []struct {
// //			name                   string
// //			getParams              string
// //			expectedStatusCode     int
// //			expectedResponseHeader string
// //			newRequest             request
// //		}{
// //			{
// //				name:                   "Ok GET",
// //				getParams:              "dkL",
// //				expectedStatusCode:     307,
// //				expectedResponseHeader: "https://google.com",
// //				newRequest:             request{inputBody: "https://google.com"},
// //			}}
// //		for _, test := range tests {
// //			t.Run(test.name, func(t *testing.T) {
// //				// Init Endpoint
// //				r := gin.New()
// //				st := inmemory.NewInMemory()
// //				cfg := config.GetConfig()
// //				r.POST("/", New(st, cfg).PostNewUserURL)
// //				// Create Request
// //				w := httptest.NewRecorder()
// //				nReq := httptest.NewRequest("POST", "/",
// //					bytes.NewBufferString(test.newRequest.inputBody))
// //
// //				// Make Request
// //				r.ServeHTTP(w, nReq)
// //				r2 := gin.New()
// //				r2.GET("/:Identifier", New(st, cfg).GetUserURLByIdentifier)
// //				// Create Request
// //				w2 := httptest.NewRecorder()
// //				req := httptest.NewRequest("GET", "/"+test.getParams, nil)
// //
// //				// Make Request
// //				r2.ServeHTTP(w2, req)
// //
// //				// Assert
// //				assert.Equal(t, test.expectedStatusCode, w2.Code)
// //				assert.Equal(t, test.expectedResponseHeader, w2.Header().Get("Location"))
// //			})
// //		}
// //	})
// //}
//
// //func TestHandler_PostNewUserURLJSON(t *testing.T) {
// //	t.Run("handler", func(t *testing.T) {
// //		tests := []struct {
// //			name                 string
// //			inputBody            string
// //			expectedStatusCode   int
// //			expectedResponseBody string
// //		}{
// //			{
// //				name:                 "Ok POST",
// //				inputBody:            `{"url": "https://google.com/new2"}`,
// //				expectedStatusCode:   201,
// //				expectedResponseBody: `{"result": "http://127.0.0.1:8080/dkM"}`,
// //			}}
// //		for _, test := range tests {
// //			t.Run(test.name, func(t *testing.T) {
// //				// Init Endpoint
// //				r := gin.New()
// //
// //				st := storage.NewInMemory()
// //
// //				r.POST("/api/shorten", New(st).PostNewUserURLJSON)
// //
// //				// Create Request
// //				m, b := map[string]string{"url": "https://google.com/new2"}, new(bytes.Buffer)
// //				json.NewEncoder(b).Encode(m)
// //
// //				w := httptest.NewRecorder()
// //				req := httptest.NewRequest("POST", "/api/shorten",
// //					b)
// //
// //				// Make Request
// //				r.ServeHTTP(w, req)
// //
// //				// Assert
// //				assert.Equal(t, test.expectedStatusCode, w.Code)
// //				assert.JSONEq(t, test.expectedResponseBody, w.Body.String())
// //			})
// //		}
// //	})
// //}
package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spaceosint/short-url/internal/config"
	"github.com/spaceosint/short-url/internal/storage/inmemory"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//	func TestGetUsersURL(t *testing.T) {
//		r := gin.New()
//
//		// Создание хранилища для теста
//		cfg := config.Config{BaseURL: "http://127.0.0.1:8080", FileStoragePath: "", ServerAddress: "127.0.0.1:8080"}
//		storage := inmemory.NewInMemory(cfg)
//		shortURL1, err := storage.GetShortURL("https://example.com")
//		if err != nil {
//			return
//		}
//		shortURL2, err := storage.GetShortURL("https://example2.net")
//		if err != nil {
//			return
//		}
//		fmt.Println(shortURL1, shortURL2)
//		// Создание тестового HTTP-запроса
//		req := httptest.NewRequest("GET", "/fwfwrfwfwhfwedscwewfgtgbrgf3r34fwc43c34fcwcxe2d2f43g544g5g34f24f23f4f", nil)
//		req.Header.Set("Accept-Encoding", "gzip")
//
//		// Создание тестового HTTP-ответа
//		w := httptest.NewRecorder()
//
//		// Вызов хендлера с тестовыми запросом и ответом
//		r.ServeHTTP(w, req)
//
//		// Проверка кода состояния ответа
//		if w.Code != http.StatusOK {
//			t.Errorf("код состояния неверный: получили %v, ожидали %v", w.Code, http.StatusOK)
//		}
//
//		// Проверка тела ответа
//		var users []map[string]string
//		err = json.Unmarshal(w.Body.Bytes(), &users)
//		if err != nil {
//			t.Errorf("ошибка декодирования тела ответа: %v", err)
//		}
//
//		expected := []map[string]string{
//			{"id": shortURL1, "url": "https://example.com"},
//			{"id": shortURL2, "url": "https://example2.net"},
//		}
//
//		if !reflect.DeepEqual(users, expected) {
//			t.Errorf("неверный список пользователей: получили %v, ожидали %v", users, expected)
//		}
//	}
func TestPostNewUserURLJSON(t *testing.T) {
	cfg := config.Config{BaseURL: "http://127.0.0.1:8080", FileStoragePath: "", ServerAddress: "127.0.0.1:8080"}

	storage := inmemory.NewInMemory(cfg)

	handler := New(storage, cfg)

	r := gin.Default()
	r.POST("/api-test", handler.PostNewUserURLJSON)

	testURL := "https://example.com"
	testBody := fmt.Sprintf(`{"original_url": "%s"}`, testURL)
	req, err := http.NewRequest("POST", "/api-test", strings.NewReader(testBody))
	if err != nil {
		t.Fatal(err)
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")

	// Perform request
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Check response status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check response body

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.JSONEq(t, `{"result":"http://127.0.0.1:8080/dkK"}`, rr.Body.String())

}
