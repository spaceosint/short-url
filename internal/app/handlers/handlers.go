package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/encoding/json"
	"github.com/spaceosint/short-url/internal/config"

	"github.com/spaceosint/short-url/internal/storage"
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
	cfg     config.Config
}

func New(storage storage.Storage, config config.Config) *Handler {
	return &Handler{
		cfg:     config,
		storage: storage,
	}
}

func (h *Handler) GetUsersURL(c *gin.Context) {
	//users, err := h.storage.GetAll()
	users, err := h.storage.GetAll(h.cfg)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}
func (h *Handler) PostNewUserURLJSON(c *gin.Context) {
	var newUserURL storage.UserURL

	if err := json.NewDecoder(c.Request.Body).Decode(&newUserURL); err != nil {
		c.IndentedJSON(http.StatusBadRequest, MyError{"bed_request"})
		return
	}

	shortURL, err := h.storage.GetShortURL(h.cfg, newUserURL.OriginalURL)
	if err != nil {
		log.Fatal(err)
	}
	newUserURL.Identifier = shortURL

	buf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false) // без этой опции символ '&' будет заменён на "\u0026"
	encoder.Encode(newUserURL.Identifier)
	fmt.Println(buf.String())

	c.IndentedJSON(http.StatusCreated, gin.H{"result": newUserURL.Identifier})
}
func (h *Handler) PostNewUserURL(c *gin.Context) {

	newUserURL, err := c.GetRawData()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, MyError{"bed_request"})
		return
	}

	shortURL, err := h.storage.GetShortURL(h.cfg, string(newUserURL))
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusCreated, shortURL)
}
func (h *Handler) GetUserURLByIdentifier(c *gin.Context) {
	id := c.Param("Identifier")
	//OriginalURL, err := h.storage.GetOriginalURL(id) storage.NewInMemory().
	OriginalURL, err := h.storage.GetOriginalURL(h.cfg, id)
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
