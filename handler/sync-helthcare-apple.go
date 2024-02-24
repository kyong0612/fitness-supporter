package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/kyong0612/fitness-supporter/infra/config"
	"github.com/kyong0612/fitness-supporter/infra/gcs"
	"github.com/kyong0612/fitness-supporter/infra/pubsub"
	handlerv1 "github.com/kyong0612/fitness-supporter/proto/generated/proto/handler/v1"
	"go.opentelemetry.io/otel"
)

// TODO: too long
//
//nolint:funlen
func (h handler) SyncHealthcareApple(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("").Start(r.Context(), "SyncHealthcareApple")
	defer span.End()

	var (
		body []byte
		err  error
	)

	if r.Body != nil {
		body, err = io.ReadAll(r.Body)
		if err != nil {
			span.RecordError(err)
			slog.ErrorContext(ctx, "failed to read request body", slog.Any("err", err))

			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	// upload sync data to gcs
	gcsClient, err := gcs.NewClient(ctx)
	if err != nil {
		span.RecordError(err)
		slog.ErrorContext(ctx, "failed to create gcs client", slog.Any("err", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	bucket := config.Get().GCSBucketFitnessSupporter
	fileName := fmt.Sprintf("%s/%s/%d.%s", "sync", "apple-healthcare", time.Now().UnixNano(), "json")

	if err := gcsClient.Upload(ctx, bucket, fileName, body); err != nil {
		span.RecordError(err)
		slog.ErrorContext(ctx, "failed to upload to gcs", slog.Any("err", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// publish RMU topic
	pubsubClient, err := pubsub.NewClient(ctx)
	if err != nil {
		span.RecordError(err)
		slog.ErrorContext(ctx, "failed to create pubsub client", slog.Any("err", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data, err := json.Marshal(handlerv1.RMUAppleHealthcareRequest{
		ObjectPath: fileName,
	})
	if err != nil {
		span.RecordError(err)
		slog.ErrorContext(ctx, "failed to marshal RMU apple healthcare request", slog.Any("err", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	topicMsg, err := pubsub.NewTopic(data, "")
	if err != nil {
		span.RecordError(err)
		slog.ErrorContext(ctx, "failed to create topic message", slog.Any("err", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := pubsubClient.PublishTopic(
		ctx,
		config.Get().PubSubTopicRMUHealthcareApple,
		topicMsg,
	); err != nil {
		span.RecordError(err)
		slog.ErrorContext(ctx, "failed to publish topic", slog.Any("err", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
