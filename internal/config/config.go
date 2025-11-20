package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database Database `yaml:"db"`
	Server   Server   `yaml:"server"`
}

type Database struct {
	DBName        string `yaml:"db_name"`
	Hostname      string `yaml:"hostname"`
	Password      string `yaml:"password"`
	Username      string `yaml:"username"`
	MigrationsDir string `yaml:"migrations_dir"`
}

type Server struct {
	Addr string `yaml:"addr"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
