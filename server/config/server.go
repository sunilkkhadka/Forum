package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func LoadServerConfig() (*ServerConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading default settings.")
	}

	address := os.Getenv("SERVER_HOST_ADDRESS")
	if address == "" {
		address = ":8081"
	}

	readTimeout, _ := time.ParseDuration(os.Getenv("SERVER_READ_TIMEOUT"))
	if readTimeout == 0 {
		readTimeout = 10 * time.Second
	}

	writeTimeout, _ := time.ParseDuration(os.Getenv("SERVER_WRITE_TIMEOUT"))
	if writeTimeout == 0 {
		writeTimeout = 30 * time.Second
	}

	idleTimeout, _ := time.ParseDuration(os.Getenv("SERVER_IDLE_TIMEOUT"))
	if idleTimeout == 0 {
		idleTimeout = time.Minute
	}

	return &ServerConfig{
		Addr:         fmt.Sprintf(":%s", address),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}, nil
}
