package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/encoding/json"
	"github.com/spaceosint/short-url/internal/config"
	"github.com/spaceosint/short-url/internal/storage/filestore"
	"github.com/spaceosint/short-url/internal/storage/inmemory"

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
	storage     inmemory.Storage
	cfg         config.Config
	fileStorage filestore.FileStore
}

func New(storage inmemory.Storage, config config.Config) *Handler {
	return &Handler{
		cfg:     config,
		storage: storage,
	}
}

type RespStruct struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

func (h *Handler) GetUserURL(c *gin.Context) {
	uuid, _ := c.Get("userID")
	if h.cfg.FileStoragePath != "" {
		userURLS := h.fileStorage.GetAllByCookieFile(h.cfg, uuid, h.cfg.FileStoragePath)
		if userURLS == nil {
			c.Status(http.StatusNoContent)
			return
		}
		c.IndentedJSON(http.StatusOK, userURLS)
	}

	userURLS, err := h.storage.GetAllByCookie(h.cfg, uuid)
	if err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.IndentedJSON(http.StatusOK, userURLS)

}
func (h *Handler) GetUsersURL(c *gin.Context) {

	if h.cfg.FileStoragePath != "" {
		users := h.fileStorage.GetAllByPathFile(h.cfg.FileStoragePath)
		c.IndentedJSON(http.StatusOK, users)
		return
	}
	users, err := h.storage.GetAll()
	fmt.Println("users")
	fmt.Println(users, err)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}
func (h *Handler) PostNewUserURLJSON(c *gin.Context) {
	var newUserURL inmemory.UserURL

	if err := json.NewDecoder(c.Request.Body).Decode(&newUserURL); err != nil {
		c.IndentedJSON(http.StatusBadRequest, MyError{"bed_request_json"})
		return
	}
	uuid, _ := c.Get("userID")
	if h.cfg.FileStoragePath != "" {

		resp, err := h.fileStorage.AddNewLinkFile(uuid, h.cfg, newUserURL.OriginalURL)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, MyError{"bed_request"})
			return
		}
		c.IndentedJSON(http.StatusCreated, gin.H{"result": resp})
		return
	}

	shortURL, err := h.storage.GetShortURL(uuid, h.cfg, newUserURL.OriginalURL)
	if err != nil {
		log.Fatal(err)
	}
	newUserURL.Identifier = shortURL

	c.IndentedJSON(http.StatusCreated, gin.H{"result": newUserURL.Identifier})
}
func (h *Handler) PostNewUserURL(c *gin.Context) {

	newUserURL, err := c.GetRawData()
	uuid, _ := c.Get("userID")

	if h.cfg.FileStoragePath != "" {

		resp, err := h.fileStorage.AddNewLinkFile(uuid, h.cfg, string(newUserURL))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, MyError{"bed_request"})
			return
		}
		c.String(http.StatusCreated, resp)
		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, MyError{"bed_request"})
		return
	}

	shortURL, err := h.storage.GetShortURL(uuid, h.cfg, string(newUserURL))
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusCreated, shortURL)
}

func (h *Handler) GetUserURLByIdentifier(c *gin.Context) {
	id := c.Param("Identifier")
	//OriginalURL, err := h.storage.GetOriginalURL(id) storage.NewInMemory().

	if h.cfg.FileStoragePath != "" {
		OriginalURL, err := h.fileStorage.GetOriginalURLFile(id, h.cfg.FileStoragePath)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, MyError{"bed_request"})
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, OriginalURL)
		return
	}

	OriginalURL, err := h.storage.GetOriginalURL(id)
	if errors.Is(err, inmemory.ErrNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "URL not found"})
		return
	}
	if errors.Is(err, inmemory.ErrAlreadyExists) {
		c.IndentedJSON(http.StatusAlreadyReported, gin.H{"message": "AlreadyExists"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, OriginalURL)

}
