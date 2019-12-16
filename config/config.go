package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	BindPort string `env:"BIND_PORT" envDefault:":8080"`
}

func Get() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, err
}
