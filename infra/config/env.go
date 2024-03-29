package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/cockroachdb/errors"
)

type config struct {
	ENV  string `env:"ENV"  envDefault:"local"`
	Port int    `env:"PORT" envDefault:"8080"`

	LINEChannelSecret      string `env:"LINE_CHANNEL_SECRET,required"`
	LINEChannelAccessToken string `env:"LINE_CHANNEL_ACCESS_TOKEN,required"`

	GCPProjectID string `env:"PROJECT_ID,required"`

	GeminiAPIKey string `env:"GEMINI_API_KEY,required"`

	GCSBucketFitnessSupporter     string `env:"GCS_BUCKET_FITNESS_SUPPORTER,required"`
	PubSubTopicAnalyzeImage       string `env:"PUBSUB_TOPIC_ANALYZE_IMAGE,required"`
	PubSubTopicRMUHealthcareApple string `env:"PUBSUB_TOPIC_RMU_HEALTHCARE_APPLE,required"`
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

func IsLocal() bool {
	return cfg.ENV == "local"
}
