// config/config.go
package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

// ServiceConfig는 개별 서비스의 설정을 나타냅니다.
type ServiceConfig struct {
	Name     string `toml:"name"`
	Endpoint string `toml:"endpoint"`
	Interval int    `toml:"interval"` // 호출 주기 (초 단위)
}

// DatabaseConfig는 데이터베이스 설정을 나타냅니다.
type DatabaseConfig struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Name     string `toml:"name"`
}

// CrawlerConfig는 전체 크롤러 설정을 나타냅니다.
type Config struct {
	General  GeneralConfig   `toml:"general"`
	Database DatabaseConfig  `toml:"database"`
	Services []ServiceConfig `toml:"services"`
}

// GeneralConfig는 일반적인 설정을 나타냅니다.
type GeneralConfig struct {
	LogLevel string `toml:"log_level"`
}

// LoadConfig는 TOML 파일에서 크롤러 설정을 읽어옵니다.
func LoadConfig(configPath string) (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found: %s", configPath)
	}

	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}
