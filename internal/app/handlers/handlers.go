package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/encoding/json"
	"github.com/spaceosint/short-url/internal/config"
	"github.com/spaceosint/short-url/internal/storage"
	"github.com/spaceosint/short-url/internal/storage/filestore"
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

type Handler struct {
	storage     storage.Storage
	cfg         config.ConfigViper
	fileStorage filestore.FileStore
}

func New(storage storage.Storage, config config.ConfigViper) *Handler {
	return &Handler{
		cfg:     config,
		storage: storage,
	}
}

func (h *Handler) GetUsersURL(c *gin.Context) {

	users, err := h.storage.GetAll()
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}
func (h *Handler) PostNewUserURLJSON(c *gin.Context) {
	var newUserURL storage.UserURL

	if err := json.NewDecoder(c.Request.Body).Decode(&newUserURL); err != nil {
		c.IndentedJSON(http.StatusBadRequest, MyError{err.Error()})
		return
	}

	shortURL, err := h.storage.GetShortURL(newUserURL.OriginalURL)
	if err != nil {
		log.Println(err)
	}
	newUserURL.Identifier = shortURL

	c.IndentedJSON(http.StatusCreated, gin.H{"result": newUserURL.Identifier})
}
func (h *Handler) PostNewUserURL(c *gin.Context) {

	newUserURL, err := c.GetRawData()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, MyError{"bed_request"})
		return
	}
	shortURL, err := h.storage.GetShortURL(string(newUserURL))
	if err != nil {
		log.Println(err)
	}
	c.String(http.StatusCreated, shortURL)
}

func (h *Handler) GetUserURLByIdentifier(c *gin.Context) {
	id := c.Param("Identifier")

	OriginalURL, err := h.storage.GetOriginalURL(id)
	if errors.Is(err, storage.ErrNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "URL not found"})
		return
	}
	if errors.Is(err, storage.ErrAlreadyExists) {
		c.IndentedJSON(http.StatusAlreadyReported, gin.H{"message": "AlreadyExists"})
		return
	}
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "bed_request"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, OriginalURL)

}
