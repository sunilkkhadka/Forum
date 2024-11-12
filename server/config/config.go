package config

import "fmt"

type Config struct {
	Server *ServerConfig
}

func LoadConfig() (*Config, error) {

	serverConfig, err := LoadServerConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load server config: %w", err)
	}

	return &Config{
		Server: serverConfig,
	}, nil
}
