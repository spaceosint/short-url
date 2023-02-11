package main

import (
	"github.com/spaceosint/short-url/internal/app/pkg/app"
	"github.com/spaceosint/short-url/internal/config"
	"log"
)

func main() {

	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.GetConfig()

	err = a.Run(cfg)
	if err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}
