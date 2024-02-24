package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/cockroachdb/errors"
)

type Message = pubsub.Message

func (c client) PublishTopic(ctx context.Context, topic string, message Message) error {
	t := c.client.Topic(topic)
	defer t.Stop()

	result := t.Publish(ctx, &message)

	_, err := result.Get(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to publish topic")
	}

	return nil
}

func NewTopic(data []byte, orderKey string) (Message, error) {
	return Message{
		Data: data,
		Attributes: map[string]string{
			"Content-Type": "application/json",
		},
		OrderingKey: orderKey,
	}, nil
}
