package pubsub

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/cockroachdb/errors"
	"github.com/kyong0612/fitness-supporter/infra/config"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/gcloud"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewMockClient(ctx context.Context) (Client, error) {
	pubsubContainer, err := gcloud.RunPubsubContainer(
		ctx,
		testcontainers.WithImage("gcr.io/google.com/cloudsdktool/cloud-sdk:367.0.0-emulators"),
		gcloud.WithProjectID(config.Get().GCPProjectID),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to run pubsub container")
	}

	projectID := pubsubContainer.Settings.ProjectID

	conn, err := grpc.Dial(pubsubContainer.URI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial pubsub container")
	}

	options := []option.ClientOption{option.WithGRPCConn(conn)}

	c, err := pubsub.NewClient(ctx, projectID, options...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create pubsub client")
	}

	// create topic (ref:https://cloud.google.com/pubsub/docs/create-topic?hl=ja#create_a_topic_2
	topic, err := c.CreateTopic(ctx, config.Get().PubSubTopicAnalyzeImage)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create topic")
	}

	// create subscription (ref:https://cloud.google.com/pubsub/docs/create-push-subscription?hl=ja#create_a_push_subscription)
	_, err = c.CreateSubscription(ctx, "fitness-supporter-test-analyze-image-event-trriger", pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 10 * time.Second,
		PushConfig: pubsub.PushConfig{
			Endpoint: fmt.Sprintf("http://localhost:%d/proto.handler.v1.AnalyzeImageService/AnalyzeImage", config.Get().Port),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create subscription")
	}

	return &client{c}, nil
}
