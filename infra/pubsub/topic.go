package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/cockroachdb/errors"
	"go.opentelemetry.io/otel"
)

type Message = pubsub.Message

func (c client) PublishTopic(ctx context.Context, topic string, message Message) error {
	ctx, span := otel.Tracer("").Start(ctx, "PublishTopic")
	defer span.End()

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
