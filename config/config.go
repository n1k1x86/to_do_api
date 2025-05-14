package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ToDoDBConfig ToDoDB `yaml:"to_do_db"`
}

type ToDoDB struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Server   string `yaml:"server"`
	DBName   string `yaml:"db_name"`
}

func (t *ToDoDB) BuildDSN() string {
	return fmt.Sprintf(`postgres://%s:%s@%s/%s?sslmode=disable`, t.Username, t.Password, t.Server, t.DBName)
}

func NewConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
