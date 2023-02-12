package main

import (
	"github.com/spaceosint/short-url/internal/app/pkg/app"
	"github.com/spaceosint/short-url/internal/config"
	"log"
)

func main() {
	cfg := config.GetConfig()

	a, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run(cfg)
	if err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}
