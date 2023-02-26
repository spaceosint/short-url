package app

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spaceosint/short-url/internal/app/handlers"
	"github.com/spaceosint/short-url/internal/config"
	"github.com/spaceosint/short-url/internal/middleware"
	"github.com/spaceosint/short-url/internal/storage/inmemory"
	"log"
)

type App struct {
	st inmemory.Storage
	h  *handlers.Handler
	r  *gin.Engine
	m  middleware.Middleware
}

func New(cfg config.Config) (*App, error) {
	a := &App{}
	s := inmemory.NewInMemory()
	a.h = handlers.New(s, cfg)
	a.r = gin.Default()

	a.r.Use(middleware.GzipHandle())
	a.r.Use(gzip.Gzip(gzip.DefaultCompression))
	a.r.Use(middleware.Authenticate())

	a.r.GET("/api/user/urls", a.h.GetUserURL)
	a.r.GET("/fwfwrfwfwhfwedscwewfgtgbrgf3r34fwc43c34fcwcxe2d2f43g544g5g34f24f23f4f", a.h.GetUsersURL)
	a.r.GET("/:Identifier", a.h.GetUserURLByIdentifier)
	a.r.POST("/", a.h.PostNewUserURL)
	a.r.POST("/api/shorten", a.h.PostNewUserURLJSON)
	a.r.GET("/ping", a.h.GetPostgreSQLPing)

	//a.r.RedirectTrailingSlash = false
	return a, nil
}

func (a *App) Run(cfg config.Config) error {

	err := a.r.Run(cfg.ServerAddress)
	if err != nil {
		log.Fatalf("Server started but then stopped error: %v", err)
	}
	return nil
}
