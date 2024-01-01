package config

import (
	"github.com/caarlos0/env/v10"
)

type config struct {
	ENV              string `env:"ENV" envDefault:"local"`
	Port             int    `env:"PORT" envDefault:"8080"`
	LINEChannelToken string `env:"LINE_CHANNEL_TOKEN,required"`
}

var (
	cfg config
)

func New() error {
	return env.Parse(&cfg)
}

func Get() config {
	return cfg
}
