package app

import (
	"ShortURL/internal/app/handlers"
	"ShortURL/internal/app/shorten"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	h *handlers.Handler
	s *shorten.Shorten
	r *gin.Engine
}

func New() (*App, error) {
	a := &App{}
	a.s = shorten.New()
	a.h = handlers.New(a.s)
	a.r = gin.Default()
	a.r.GET("/fwfwrfwfwhfwedscwewfgtgbrgf3r34fwc43c34fcwcxe2d2f43g544g5g34f24f23f4f", a.h.GetUsersURL)
	a.r.GET("/:Identifier", a.h.GetUserURLByIdentifier)
	a.r.POST("/", a.h.PostNewUserURL)
	//a.r.RedirectTrailingSlash = false
	return a, nil
}

func (a *App) Run() error {
	err := a.r.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Server started but then stopped error: %v", err)
	}
	return nil
}
