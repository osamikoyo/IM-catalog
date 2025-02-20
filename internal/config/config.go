package config

import (
	"fmt"
	"io"
	"os"

	"github.com/bytedance/sonic"
)

type Config struct{
	Port string `yaml:"pory"`
	Host string `yaml:"host"`
    DatabaseURL string `yaml:"dsn"`
}

func Load(cfgPath string) (*Config, error) {
	file, err := os.Open(cfgPath)
	if err != nil{
		return nil, fmt.Errorf("cant open config file: %v", err)
	}

	body, err := io.ReadAll(file)
	if err != nil{
		return nil, fmt.Errorf("cant read file: %v", err)
	}

	var cfg Config

	err = sonic.Unmarshal(body, &cfg)
	return &cfg, err
} 
