package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct{
	Port uint32 `toml:"port"`
	Host string `toml:"host"`
    DatabaseURL string `toml:"dsn"`
}

func Load(cfgPath string) (*Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(cfgPath, &cfg)
	return &cfg, err
} 
