package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ConfigViper struct {
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
}

func GetConfigViper() ConfigViper {
	// Инициализация флагов
	serverAddress := pflag.StringP("SERVER_ADDRESS", "a", "", "Server address")
	baseURL := pflag.StringP("BASE_URL", "b", "", "Base URL")
	fileStoragePath := pflag.StringP("FILE_STORAGE_PATH", "f", "", "File storage path")

	// Сообщаем флагам о необходимости парсинга
	pflag.Parse()

	// Инициализация Viper
	viper.AutomaticEnv()        // автоматическое чтение переменных окружения
	viper.SetConfigType("yaml") // тип конфигурации

	// Чтение значений из env
	viper.SetDefault("SERVER_ADDRESS", "localhost:8080")
	viper.SetDefault("BASE_URL", "http://localhost:8080")
	viper.SetDefault("FILE_STORAGE_PATH", "")

	// Применение значений из флагов, если они были установлены
	if *serverAddress != "" {
		viper.Set("SERVER_ADDRESS", *serverAddress)
	}
	if *baseURL != "" {
		viper.Set("BASE_URL", *baseURL)
	}
	if *fileStoragePath != "" {
		viper.Set("FILE_STORAGE_PATH", *fileStoragePath)
	}

	return ConfigViper{
		ServerAddress:   viper.GetString("SERVER_ADDRESS"),
		BaseURL:         viper.GetString("BASE_URL"),
		FileStoragePath: viper.GetString("FILE_STORAGE_PATH"),
	}
}
