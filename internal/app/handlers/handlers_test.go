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

func TestPostNewUserURLJSON(t *testing.T) {
	cfg := config.ConfigViper{BaseURL: "http://127.0.0.1:8080", FileStoragePath: "", ServerAddress: "127.0.0.1:8080"}

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
