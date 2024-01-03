package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/cockroachdb/errors"
)

type config struct {
	ENV  string `env:"ENV"                                envDefault:"local"`
	Port int    `env:"PORT"                               envDefault:"8080"`

	LINEChannelSecret      string `env:"LINE_CHANNEL_SECRET,required"`
	LINEChannelAccessToken string `env:"LINE_CHANNEL_ACCESS_TOKEN,required"`

	GeminiAPIKey string `env:"GEMINI_API_KEY,required"`

	GCSBucketFitnessSupporter string `env:"GCS_BUCKET_FITNESS_SUPPORTER,required"`
}

var cfg config

func New() error {
	if err := env.Parse(&cfg); err != nil {
		return errors.Wrap(err, "failed to parse config")
	}

	return nil
}

func Get() config {
	return cfg
}
