package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ServerConfig struct {
	BindAddress string `yaml:"bind_address"`
	BindPort    int    `yaml:"bind_port"`
}

type Config struct {
	ServerConfig ServerConfig `yaml:"server_config"`
}

func defaultServerConfig() ServerConfig {
	return ServerConfig{
		BindAddress: "localhost",
		BindPort:    3000,
	}
}

func defaultConfig() Config {
	return Config{
		ServerConfig: defaultServerConfig(),
	}
}

func ReadConfig() (config Config, err error) {
	config = defaultConfig()
	file, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(file, &config)
	return
}
