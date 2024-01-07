package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/cockroachdb/errors"
	"github.com/kyong0612/fitness-supporter/infra/config"
)

type Client interface {
	PublishTopic(ctx context.Context, topic string, message Message) error
}

type client struct {
	client *pubsub.Client
}

func NewClient(ctx context.Context) (Client, error) {
	c, err := pubsub.NewClient(ctx, config.Get().GCPProjectID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create pubsub client")
	}

	return &client{c}, nil
}
