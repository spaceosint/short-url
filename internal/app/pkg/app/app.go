package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spaceosint/short-url/internal/app/handlers"
	"github.com/spaceosint/short-url/internal/config"
	"github.com/spaceosint/short-url/internal/middleware"
	"github.com/spaceosint/short-url/internal/storage"
	"github.com/spaceosint/short-url/internal/storage/filestore"
	"github.com/spaceosint/short-url/internal/storage/inmemory"
	"log"
)

type App struct {
	st  storage.Storage
	h   *handlers.Handler
	r   *gin.Engine
	cfg config.Config
	m   middleware.Middleware
}

func New(cfg config.Config) (*App, error) {
	a := &App{}

	if cfg.FileStoragePath != "" {
		fs := filestore.NewInFileStore(cfg)
		a.h = handlers.New(fs, cfg)
	} else {
		ms := inmemory.NewInMemory(cfg)
		a.h = handlers.New(ms, cfg)
	}

	a.r = gin.Default()

	a.r.Use(middleware.GzipReaderHandle())
	a.r.Use(middleware.GzipWriterHandler())

	a.r.GET("/fwfwrfwfwhfwedscwewfgtgbrgf3r34fwc43c34fcwcxe2d2f43g544g5g34f24f23f4f", a.h.GetUsersURL)
	a.r.GET("/:Identifier", a.h.GetUserURLByIdentifier)
	a.r.POST("/", a.h.PostNewUserURL)
	a.r.POST("/api/shorten", a.h.PostNewUserURLJSON)

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
