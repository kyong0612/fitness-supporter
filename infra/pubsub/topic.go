package pubsub

import (
	"context"
	"encoding/json"
	"log/slog"

	"cloud.google.com/go/pubsub"
	"github.com/cockroachdb/errors"
	analyzeimagev1 "github.com/kyong0612/fitness-supporter/proto/generated/proto/handler/v1"
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

func NewAnalyzeImageTopic(userID, imageURL string) (Message, error) {
	data, err := json.Marshal(analyzeimagev1.AnalyzeImageRequest{
		UserId:   userID,
		ImageUrl: imageURL,
	})
	if err != nil {
		return Message{}, errors.Wrap(err, "failed to marshal analyze image request")
	}

	// TODO: remove
	slog.Info("marshaled data",
		slog.String("data", string(data)),
	)

	return Message{
		Data: data,
		Attributes: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
