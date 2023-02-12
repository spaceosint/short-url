package main

import (
	"flag"
	"github.com/spaceosint/short-url/internal/app/pkg/app"
	"github.com/spaceosint/short-url/internal/config"
	"log"
	"os"
)

func main() {
	cfg := config.GetConfig()

	serverAddress := flag.String("a", os.Getenv("SERVER_ADDRESS"), "a string")
	fileStoragePath := flag.String("b", os.Getenv("BASE_URL"), "a string")
	baseURL := flag.String("f", os.Getenv("FILE_STORAGE_PATH"), "a string")
	flag.Parse()
	if *serverAddress != "" {
		cfg.ServerAddress = *serverAddress
	}
	if *fileStoragePath != "" {
		cfg.FileStoragePath = *fileStoragePath
	}
	if *baseURL != "" {
		cfg.BaseURL = *baseURL
	}

	a, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run(cfg)
	if err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}
