package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env          string `env:"ENV" envDefault:"dev"`
	Port         string `env:"PORT" envDefault:"80"`
	Database_url string `env:"DATABASE_URL" envDefult:""`
	ProjectID    string `env:"PROJECTID" envDefault:""`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
