package main

import (
	"fmt"
	"github.com/spaceosint/short-url/internal/app/pkg/app"
	"github.com/spaceosint/short-url/internal/config"
	"log"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	//serverAddress := flag.String("a", getEnv("SERVER_ADDRESS", "127.0.0.1:8080"), "a string")
	//baseURL := flag.String("b", getEnv("BASE_URL", "http://127.0.0.1:8080"), "a string")
	//fileStoragePath := flag.String("f", getEnv("FILE_STORAGE_PATH", ""), "a string")
	//flag.Parse()

	cfg := config.GetConfigViper()
	fmt.Println("Hello world", cfg.FileStoragePath, cfg.BaseURL, cfg.ServerAddress, cfg)
	//if cfg.ServerAddress == "127.0.0.1:8080" {
	//	cfg.ServerAddress = *serverAddress
	//}
	//if cfg.FileStoragePath == "" {
	//	cfg.FileStoragePath = *fileStoragePath
	//}
	//if cfg.BaseURL == "http://127.0.0.1:8080" {
	//	cfg.BaseURL = *baseURL
	//}

	fmt.Println(cfg)
	a, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run(cfg)
	if err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}
