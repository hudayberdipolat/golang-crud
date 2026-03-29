package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbConfig     dbConfig
	ServerConfig serverConfig
	FolderConfig folderConfig
}

type dbConfig struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbSslMode  string
	DbTimeZone string
}

type serverConfig struct {
	ServerHost string
	ServerPort string
	AppName    string
	AppHeader  string
}

type folderConfig struct {
	RootPath   string
	PublicPath string
}

func GetConfig() (*Config, error) {
	if err := godotenv.Load("./../.env"); err != nil {
		return nil, err
	}
	return &Config{
		DbConfig: dbConfig{
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbSslMode:  os.Getenv("DB_SSL_MODE"),
			DbTimeZone: os.Getenv("DB_TIME_ZONE"),
		},
		ServerConfig: serverConfig{
			ServerHost: os.Getenv("SERVER_HOST"),
			ServerPort: os.Getenv("SERVER_PORT"),
			AppName:    os.Getenv("APP_NAME"),
			AppHeader:  os.Getenv("APP_HEADER"),
		},
		FolderConfig: folderConfig{
			RootPath:   os.Getenv("ROOT_PATH"),
			PublicPath: os.Getenv("PUBLIC_PATH"),
		},
	}, nil
}
