package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/cockroachdb/errors"
	"github.com/kyong0612/fitness-supporter/infra/config"
)

type Client interface{}

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

func (c client) PublishTopic(ctx context.Context, topic string, data []byte) error {
	t := c.client.Topic(topic)
	defer t.Stop()

	result := t.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	_, err := result.Get(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to publish topic")
	}

	return nil
}
