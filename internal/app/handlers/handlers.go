package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Shorten interface {
	ShortenURL(id uint32) string
}

type Handler struct {
	s Shorten
}

func New(s Shorten) *Handler {
	return &Handler{
		s: s,
	}
}

type MyError struct {
	Error string `json:"error"`
}
type UserURL struct {
	ID          uint32
	Identifier  string
	OriginalURL string
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

var UsersURL = []UserURL{
	{ID: 1000, OriginalURL: "https://yandex.ru", Identifier: "t1"},
	{ID: 1001, OriginalURL: "https://yandex.ru/123", Identifier: "t2"},
}

func (h *Handler) GetUsersURL(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, UsersURL)
}
func (h *Handler) PostNewUserURL(c *gin.Context) {

	newUserURL, err := c.GetRawData()

	fmt.Printf("new user = [%s]", newUserURL)
	if err != nil {
		log.Print(err)
		c.IndentedJSON(http.StatusBadRequest, MyError{"bed_request"})
		return
	}

	id := getNewID()
	shortURL := h.s.ShortenURL(id)

	var newUser = UserURL{ID: id, OriginalURL: string(newUserURL), Identifier: shortURL}
	UsersURL = append(UsersURL, newUser)
	c.String(http.StatusCreated, "http://127.0.0.1:8080/"+shortURL)
}
func (h *Handler) GetUserURLByIdentifier(c *gin.Context) {
	id := c.Param("Identifier")
	for _, user := range UsersURL {
		if user.Identifier == id {
			//c.Redirect(http.StatusOK, user.OriginalURL)
			//c.Header("Location", user.OriginalURL)
			//c.Status(http.StatusTemporaryRedirect)

			c.Redirect(http.StatusTemporaryRedirect, user.OriginalURL)
			//c.IndentedJSON(http.StatusTemporaryRedirect, user.OriginalURL)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "URL not found"})
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
