package main

import (
	"flag"
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
	cfg := config.GetConfig()

	serverAddress := flag.String("a", getEnv("SERVER_ADDRESS", "127.0.0.1:8080"), "a string")
	fileStoragePath := flag.String("b", getEnv("BASE_URL", "file"), "a string")
	baseURL := flag.String("f", getEnv("FILE_STORAGE_PATH", "http://127.0.0.1:8080"), "a string")
	flag.Parse()
	//if *serverAddress != "" {
	cfg.ServerAddress = *serverAddress
	//}
	//if *fileStoragePath != "" {
	cfg.FileStoragePath = *fileStoragePath
	//}
	//if *baseURL != "" {
	cfg.BaseURL = *baseURL
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
