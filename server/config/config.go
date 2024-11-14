package config

import "fmt"

type Config struct {
	Server   *ServerConfig
	Database *DBConfig
}

func LoadConfig() (*Config, error) {

	serverConfig, err := LoadServerConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load server config: %w", err)
	}

	databaseConfig, err := LoadDBConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load database config: %w", err)
	}

	return &Config{
		Server:   serverConfig,
		Database: databaseConfig,
	}, nil
}
