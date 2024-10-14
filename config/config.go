package config

import (
	"log"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

type Config struct {
	ConnStrDb      string
	JwtSecret      string
	LogFilePath    string
	CacheAddress   string
	CachePassword  string
	CacheName      int
	MigrationsPath string
}

var Variables *Config

func Init() error {
	err := gotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		return err
	}

	cacheName, err := strconv.Atoi(os.Getenv("CACHE_DB_NAME"))
	if err != nil {
		return err
	}

	Variables = &Config{
		ConnStrDb:      os.Getenv("CONNECTION_STRING"),
		JwtSecret:      os.Getenv("JWT_SECRET"),
		LogFilePath:    os.Getenv("LOG_FILE_PATH"),
		CacheAddress:   os.Getenv("CACHE_DB_ADDRESS"),
		CachePassword:  os.Getenv("CACHE_DB_PASSWORD"),
		CacheName:      cacheName,
		MigrationsPath: os.Getenv("MIGRATIONS_PATH"),
	}

	return nil
}
