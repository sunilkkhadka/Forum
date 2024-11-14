package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	Name            string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func LoadDBConfig() (*DBConfig, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading default settings.")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return nil, fmt.Errorf("DB HOST not found")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, fmt.Errorf("DB USER not found")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return nil, fmt.Errorf("DB PASSWORD not found")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, fmt.Errorf("DB NAME not found")
	}

	dbMaxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		return nil, fmt.Errorf("DB_MAX_OPEN_CONNS should be an integer")
	}

	dbMaxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		return nil, fmt.Errorf("DB_MAX_IDLE_CONNS should be an integer")
	}

	dbConnMaxLifetime, err := time.ParseDuration(os.Getenv("DB_CONN_MAX_LIFE"))
	if err != nil {
		return nil, fmt.Errorf("DB_CONN_MAX_LIFE should be a valid duration with m/s/h for minute/seconds/hour")
	}

	return &DBConfig{
		Host:            dbHost,
		Port:            dbPort,
		User:            dbUser,
		Name:            dbName,
		Password:        dbPassword,
		MaxOpenConns:    dbMaxOpenConns,
		MaxIdleConns:    dbMaxIdleConns,
		ConnMaxLifetime: dbConnMaxLifetime,
	}, nil
}
