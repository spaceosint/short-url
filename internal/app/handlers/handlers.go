package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/spaceosint/short-url/internal/storage"
	"io"
	"log"
	"net/http"
)

type MyError struct {
	Error string `json:"error"`
}
type handler interface {
	GetUsersURL(c *gin.Context)
	PostNewUserURL(c *gin.Context)
	GetUserURLByIdentifier(c *gin.Context)
}

//type Shorten interface {
//	ShortenURL(id uint32) string
//}

type Handler struct {
	storage storage.Storage
}

func New(storage storage.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) GetUsersURL(c *gin.Context) {
	//users, err := h.storage.GetAll()
	users, err := h.storage.GetAll()
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}
func (h *Handler) PostNewUserURLJSON(c *gin.Context) {
	var newUserURL storage.UserURL
	//
	//if err := c.BindJSON(&newUserURL); err != nil {
	//	c.IndentedJSON(http.StatusBadRequest, MyError{"bed_request"})
	//	return
	//}

	b, err := io.ReadAll(c.Request.Body)
	// обрабатываем ошибку
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, MyError{"bed_request"})
		return
	}

	//ewUserURL, err := c.GetRawData()
	//if err != nil {
	//	c.IndentedJSON(http.StatusBadRequest, MyError{"bed_request"})
	//	return
	//}
	if err := json.Unmarshal(b, &newUserURL); err != nil {
		panic(err)
	}
	fmt.Println(newUserURL.OriginalURL)
	shortURL, err := h.storage.GetShortURL(newUserURL.OriginalURL)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"result": "http://127.0.0.1:8080/" + shortURL})
}
func (h *Handler) PostNewUserURL(c *gin.Context) {

	newUserURL, err := c.GetRawData()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, MyError{"bed_request"})
		return
	}

	//shortURL := h.storage.GetShortURL(string(newUserURL))
	shortURL, err := h.storage.GetShortURL(string(newUserURL))
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusCreated, "http://127.0.0.1:8080/"+shortURL)
}
func (h *Handler) GetUserURLByIdentifier(c *gin.Context) {
	id := c.Param("Identifier")
	//OriginalURL, err := h.storage.GetOriginalURL(id) storage.NewInMemory().
	OriginalURL, err := h.storage.GetOriginalURL(id)
	if errors.Is(err, storage.ErrNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "URL not found"})
		return
	}
	if errors.Is(err, storage.ErrAlreadyExists) {
		c.IndentedJSON(http.StatusAlreadyReported, gin.H{"message": "AlreadyExists"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, OriginalURL)

}
