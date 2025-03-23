package config

import (
	"github.com/pelletier/go-toml"
	"log"
	"os"
)

type Config struct {
	Server   ServerConfig   `toml:"server"`
	Database DatabaseConfig `toml:"database"`
}

type ServerConfig struct {
	ApiPort     int `toml:"api_port"`
	CrawlerPort int `toml:"crawler_port"`
}

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"name"`
}

var AppConfig *Config

// LoadConfig - Loads configuration from a TOML file
func LoadConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Parse TOML file
	decoder := toml.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return err
	}

	AppConfig = config
	return nil
}

// GetConfig - Returns the loaded configuration
func GetConfig() *Config {
	if AppConfig == nil {
		log.Fatal("Configuration not loaded. Call LoadConfig first.")
	}
	return AppConfig
}
