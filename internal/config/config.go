package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" `    // envDefault:"127.0.0.1:8080"
	BaseURL         string `env:"BASE_URL"`           // envDefault:"http://127.0.0.1:8080"
	FileStoragePath string `env:"FILE_STORAGE_PATH" ` // envDefault:"file"
}

//var instance *Config
//var once sync.Once

func GetConfig() Config {
	//once.Do(func() {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	//})
	return cfg
}
